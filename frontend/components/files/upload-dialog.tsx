"use client"

import type React from "react"

import { useState, useCallback } from "react"
import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Progress } from "@/components/ui/progress"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Upload, X, File, CheckCircle, AlertCircle } from "lucide-react"
import { useDropzone } from "react-dropzone"

interface UploadDialogProps {
  folderId: string
  onUploadComplete: () => void
  trigger?: React.ReactNode
}

interface UploadFile {
  file: File
  progress: number
  status: "pending" | "uploading" | "completed" | "error"
  error?: string
}

export function UploadDialog({ folderId, onUploadComplete, trigger }: UploadDialogProps) {
  const [isOpen, setIsOpen] = useState(false)
  const [uploadFiles, setUploadFiles] = useState<UploadFile[]>([])
  const [isUploading, setIsUploading] = useState(false)

  const onDrop = useCallback((acceptedFiles: File[]) => {
    const newFiles = acceptedFiles.map((file) => ({
      file,
      progress: 0,
      status: "pending" as const,
    }))
    setUploadFiles((prev) => [...prev, ...newFiles])
  }, [])

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    multiple: true,
  })

  const removeFile = (index: number) => {
    setUploadFiles((prev) => prev.filter((_, i) => i !== index))
  }

  const uploadFile = async (uploadFile: UploadFile, index: number) => {
    const formData = new FormData()
    formData.append("file", uploadFile.file)

    try {
      setUploadFiles((prev) => prev.map((f, i) => (i === index ? { ...f, status: "uploading" } : f)))

      const xhr = new XMLHttpRequest()

      xhr.upload.addEventListener("progress", (event) => {
        if (event.lengthComputable) {
          const progress = Math.round((event.loaded / event.total) * 100)
          setUploadFiles((prev) => prev.map((f, i) => (i === index ? { ...f, progress } : f)))
        }
      })

      xhr.addEventListener("load", () => {
        if (xhr.status === 200) {
          setUploadFiles((prev) => prev.map((f, i) => (i === index ? { ...f, status: "completed", progress: 100 } : f)))
        } else {
          setUploadFiles((prev) =>
            prev.map((f, i) => (i === index ? { ...f, status: "error", error: "アップロードに失敗しました" } : f)),
          )
        }
      })

      xhr.addEventListener("error", () => {
        setUploadFiles((prev) =>
          prev.map((f, i) => (i === index ? { ...f, status: "error", error: "ネットワークエラーが発生しました" } : f)),
        )
      })

      xhr.open("POST", `/s/${folderId}/file/upload`)
      xhr.send(formData)
    } catch (error) {
      setUploadFiles((prev) =>
        prev.map((f, i) => (i === index ? { ...f, status: "error", error: "アップロードに失敗しました" } : f)),
      )
    }
  }

  const startUpload = async () => {
    setIsUploading(true)

    const pendingFiles = uploadFiles.filter((f) => f.status === "pending")

    for (let i = 0; i < uploadFiles.length; i++) {
      if (uploadFiles[i].status === "pending") {
        await uploadFile(uploadFiles[i], i)
      }
    }

    setIsUploading(false)

    // Check if all uploads completed successfully
    const allCompleted = uploadFiles.every((f) => f.status === "completed")
    if (allCompleted) {
      onUploadComplete()
      setTimeout(() => {
        setIsOpen(false)
        setUploadFiles([])
      }, 1000)
    }
  }

  const formatFileSize = (bytes: number) => {
    const sizes = ["B", "KB", "MB", "GB"]
    const i = Math.floor(Math.log(bytes) / Math.log(1024))
    return `${(bytes / Math.pow(1024, i)).toFixed(1)} ${sizes[i]}`
  }

  const getStatusIcon = (status: UploadFile["status"]) => {
    switch (status) {
      case "completed":
        return <CheckCircle className="h-4 w-4 text-green-500" />
      case "error":
        return <AlertCircle className="h-4 w-4 text-red-500" />
      default:
        return <File className="h-4 w-4 text-muted-foreground" />
    }
  }

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        {trigger || (
          <Button>
            <Upload className="h-4 w-4 mr-2" />
            ファイルをアップロード
          </Button>
        )}
      </DialogTrigger>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>ファイルアップロード</DialogTitle>
          <DialogDescription>ファイルをドラッグ&ドロップするか、クリックして選択してください</DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          {/* Drop Zone */}
          <div
            {...getRootProps()}
            className={`border-2 border-dashed rounded-lg p-8 text-center cursor-pointer transition-colors ${
              isDragActive ? "border-primary bg-primary/5" : "border-muted-foreground/25 hover:border-primary/50"
            }`}
          >
            <input {...getInputProps()} />
            <Upload className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
            {isDragActive ? (
              <p className="text-primary">ファイルをここにドロップしてください</p>
            ) : (
              <div>
                <p className="text-foreground font-medium mb-2">ファイルをドラッグ&ドロップ</p>
                <p className="text-muted-foreground text-sm">
                  または <span className="text-primary">クリックして選択</span>
                </p>
              </div>
            )}
          </div>

          {/* File List */}
          {uploadFiles.length > 0 && (
            <div className="space-y-2 max-h-60 overflow-y-auto">
              {uploadFiles.map((uploadFile, index) => (
                <div key={index} className="flex items-center gap-3 p-3 border rounded-lg">
                  {getStatusIcon(uploadFile.status)}

                  <div className="flex-1 min-w-0">
                    <p className="text-sm font-medium truncate">{uploadFile.file.name}</p>
                    <p className="text-xs text-muted-foreground">{formatFileSize(uploadFile.file.size)}</p>

                    {uploadFile.status === "uploading" && <Progress value={uploadFile.progress} className="mt-2" />}

                    {uploadFile.status === "error" && uploadFile.error && (
                      <Alert variant="destructive" className="mt-2">
                        <AlertDescription className="text-xs">{uploadFile.error}</AlertDescription>
                      </Alert>
                    )}
                  </div>

                  {uploadFile.status === "pending" && (
                    <Button variant="ghost" size="sm" onClick={() => removeFile(index)}>
                      <X className="h-4 w-4" />
                    </Button>
                  )}
                </div>
              ))}
            </div>
          )}

          {/* Actions */}
          <div className="flex justify-end gap-2">
            <Button
              variant="outline"
              onClick={() => {
                setIsOpen(false)
                setUploadFiles([])
              }}
              disabled={isUploading}
            >
              キャンセル
            </Button>
            <Button onClick={startUpload} disabled={uploadFiles.length === 0 || isUploading}>
              {isUploading ? "アップロード中..." : "アップロード開始"}
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}
