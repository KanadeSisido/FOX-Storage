"use client"

import type React from "react"

import { useState } from "react"
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
  ContextMenuSeparator,
} from "@/components/ui/context-menu"
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Download, Edit, Trash2, Copy } from "lucide-react"

interface FileItem {
  id: string
  name: string
  type: "file" | "folder"
  size?: number
  createdAt: string
  mimeType?: string
}

interface FileContextMenuProps {
  file: FileItem
  children: React.ReactNode
  onDownload?: (fileId: string, fileName: string) => void
  onDelete?: (fileId: string) => void
  onRename?: (fileId: string, newName: string) => void
}

export function FileContextMenu({ file, children, onDownload, onDelete, onRename }: FileContextMenuProps) {
  const [showRenameDialog, setShowRenameDialog] = useState(false)
  const [showDeleteDialog, setShowDeleteDialog] = useState(false)
  const [newName, setNewName] = useState(file.name)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState("")

  const handleRename = async () => {
    if (!newName.trim() || newName === file.name) {
      setShowRenameDialog(false)
      return
    }

    setIsLoading(true)
    setError("")

    try {
      if (onRename) {
        await onRename(file.id, newName.trim())
        setShowRenameDialog(false)
      }
    } catch (err) {
      setError("名前の変更に失敗しました")
    } finally {
      setIsLoading(false)
    }
  }

  const handleDelete = async () => {
    setIsLoading(true)
    setError("")

    try {
      if (onDelete) {
        await onDelete(file.id)
        setShowDeleteDialog(false)
      }
    } catch (err) {
      setError("削除に失敗しました")
    } finally {
      setIsLoading(false)
    }
  }

  const copyToClipboard = async () => {
    try {
      await navigator.clipboard.writeText(file.name)
    } catch (err) {
      console.error("Failed to copy to clipboard:", err)
    }
  }

  return (
    <>
      <ContextMenu>
        <ContextMenuTrigger>{children}</ContextMenuTrigger>
        <ContextMenuContent>
          {file.type === "file" && onDownload && (
            <>
              <ContextMenuItem onClick={() => onDownload(file.id, file.name)}>
                <Download className="h-4 w-4 mr-2" />
                ダウンロード
              </ContextMenuItem>
              <ContextMenuSeparator />
            </>
          )}

          <ContextMenuItem onClick={() => setShowRenameDialog(true)}>
            <Edit className="h-4 w-4 mr-2" />
            名前を変更
          </ContextMenuItem>

          <ContextMenuItem onClick={copyToClipboard}>
            <Copy className="h-4 w-4 mr-2" />
            名前をコピー
          </ContextMenuItem>

          <ContextMenuSeparator />

          <ContextMenuItem
            onClick={() => setShowDeleteDialog(true)}
            className="text-destructive focus:text-destructive"
          >
            <Trash2 className="h-4 w-4 mr-2" />
            削除
          </ContextMenuItem>
        </ContextMenuContent>
      </ContextMenu>

      {/* Rename Dialog */}
      <Dialog open={showRenameDialog} onOpenChange={setShowRenameDialog}>
        <DialogContent className="max-w-md">
          <DialogHeader>
            <DialogTitle>名前を変更</DialogTitle>
            <DialogDescription>新しい名前を入力してください</DialogDescription>
          </DialogHeader>

          <div className="space-y-4">
            {error && (
              <Alert variant="destructive">
                <AlertDescription>{error}</AlertDescription>
              </Alert>
            )}

            <div className="space-y-2">
              <Label htmlFor="newName">新しい名前</Label>
              <Input
                id="newName"
                type="text"
                value={newName}
                onChange={(e) => setNewName(e.target.value)}
                onKeyDown={(e) => {
                  if (e.key === "Enter") {
                    handleRename()
                  }
                }}
              />
            </div>

            <div className="flex justify-end gap-2">
              <Button
                variant="outline"
                onClick={() => {
                  setShowRenameDialog(false)
                  setNewName(file.name)
                  setError("")
                }}
                disabled={isLoading}
              >
                キャンセル
              </Button>
              <Button onClick={handleRename} disabled={isLoading}>
                {isLoading ? "変更中..." : "変更"}
              </Button>
            </div>
          </div>
        </DialogContent>
      </Dialog>

      {/* Delete Dialog */}
      <Dialog open={showDeleteDialog} onOpenChange={setShowDeleteDialog}>
        <DialogContent className="max-w-md">
          <DialogHeader>
            <DialogTitle>削除の確認</DialogTitle>
            <DialogDescription>「{file.name}」を削除してもよろしいですか？この操作は取り消せません。</DialogDescription>
          </DialogHeader>

          <div className="space-y-4">
            {error && (
              <Alert variant="destructive">
                <AlertDescription>{error}</AlertDescription>
              </Alert>
            )}

            <div className="flex justify-end gap-2">
              <Button
                variant="outline"
                onClick={() => {
                  setShowDeleteDialog(false)
                  setError("")
                }}
                disabled={isLoading}
              >
                キャンセル
              </Button>
              <Button variant="destructive" onClick={handleDelete} disabled={isLoading}>
                {isLoading ? "削除中..." : "削除"}
              </Button>
            </div>
          </div>
        </DialogContent>
      </Dialog>
    </>
  )
}
