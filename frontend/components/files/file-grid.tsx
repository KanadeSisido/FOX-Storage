"use client"

import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { FileContextMenu } from "@/components/files/file-context-menu"
import { File, Folder, Download, FileText, ImageIcon, Video, Music, Archive, FileCode } from "lucide-react"

interface FileItem {
  id: string
  name: string
  type: "file" | "folder"
  size?: number
  createdAt: string
  mimeType?: string
}

interface FileGridProps {
  files: FileItem[]
  onFileClick: (file: FileItem) => void
  onDownload: (fileId: string, fileName: string) => void
  onDelete?: (fileId: string) => void
  onRename?: (fileId: string, newName: string) => void
}

export function FileGrid({ files, onFileClick, onDownload, onDelete, onRename }: FileGridProps) {
  const getFileIcon = (file: FileItem) => {
    if (file.type === "folder") {
      return <Folder className="h-8 w-8 text-primary" />
    }

    const mimeType = file.mimeType || ""
    if (mimeType.startsWith("image/")) {
      return <ImageIcon className="h-8 w-8 text-green-500" />
    }
    if (mimeType.startsWith("video/")) {
      return <Video className="h-8 w-8 text-red-500" />
    }
    if (mimeType.startsWith("audio/")) {
      return <Music className="h-8 w-8 text-purple-500" />
    }
    if (mimeType.includes("text/") || mimeType.includes("document")) {
      return <FileText className="h-8 w-8 text-blue-500" />
    }
    if (mimeType.includes("zip") || mimeType.includes("archive")) {
      return <Archive className="h-8 w-8 text-orange-500" />
    }
    if (mimeType.includes("code") || mimeType.includes("javascript") || mimeType.includes("json")) {
      return <FileCode className="h-8 w-8 text-indigo-500" />
    }

    return <File className="h-8 w-8 text-muted-foreground" />
  }

  const formatFileSize = (bytes?: number) => {
    if (!bytes) return ""

    const sizes = ["B", "KB", "MB", "GB"]
    const i = Math.floor(Math.log(bytes) / Math.log(1024))
    return `${(bytes / Math.pow(1024, i)).toFixed(1)} ${sizes[i]}`
  }

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("ja-JP", {
      year: "numeric",
      month: "short",
      day: "numeric",
    })
  }

  if (files.length === 0) {
    return (
      <div className="text-center py-12">
        <Folder className="h-16 w-16 text-muted-foreground mx-auto mb-4" />
        <h3 className="text-lg font-medium text-foreground mb-2">フォルダが空です</h3>
        <p className="text-muted-foreground">ファイルをアップロードするか、新しいフォルダを作成してください。</p>
      </div>
    )
  }

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
      {files.map((file) => (
        <FileContextMenu
          key={file.id}
          file={file}
          onDownload={file.type === "file" ? onDownload : undefined}
          onDelete={onDelete}
          onRename={onRename}
        >
          <Card className="hover:shadow-md transition-shadow cursor-pointer group" onClick={() => onFileClick(file)}>
            <CardContent className="p-4">
              <div className="flex flex-col items-center text-center space-y-3">
                {getFileIcon(file)}

                <div className="w-full">
                  <h3 className="font-medium text-sm text-foreground truncate" title={file.name}>
                    {file.name}
                  </h3>

                  <div className="flex flex-col gap-1 mt-2">
                    {file.type === "file" && file.size && (
                      <Badge variant="secondary" className="text-xs">
                        {formatFileSize(file.size)}
                      </Badge>
                    )}
                    <p className="text-xs text-muted-foreground">{formatDate(file.createdAt)}</p>
                  </div>
                </div>

                {file.type === "file" && (
                  <Button
                    variant="outline"
                    size="sm"
                    className="w-full opacity-0 group-hover:opacity-100 transition-opacity bg-transparent"
                    onClick={(e) => {
                      e.stopPropagation()
                      onDownload(file.id, file.name)
                    }}
                  >
                    <Download className="h-4 w-4 mr-2" />
                    ダウンロード
                  </Button>
                )}
              </div>
            </CardContent>
          </Card>
        </FileContextMenu>
      ))}
    </div>
  )
}
