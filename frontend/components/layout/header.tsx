"use client"

import { Button } from "@/components/ui/button"
import { Cloud, Upload, Settings, LogOut } from "lucide-react"
import { UploadDialog } from "@/components/files/upload-dialog"
import Link from "next/link"
import { useParams } from "next/navigation"

interface HeaderProps {
  showNavigation?: boolean
  onLogout?: () => void
  onUploadComplete?: () => void
}

export function Header({ showNavigation = false, onLogout, onUploadComplete }: HeaderProps) {
  const params = useParams()
  const folderId = params?.folderId as string

  return (
    <header className="border-b border-border bg-background">
      <div className="flex h-16 items-center justify-between px-6">
        <div className="flex items-center gap-2">
          <Cloud className="h-8 w-8 text-primary" />
          <h1 className="text-xl font-bold text-foreground">CloudStorage</h1>
        </div>

        {showNavigation && (
          <nav className="flex items-center gap-4">
            <Button variant="ghost" size="sm" asChild>
              <Link href="/dashboard">
                <Cloud className="h-4 w-4 mr-2" />
                ホーム
              </Link>
            </Button>

            <UploadDialog
              folderId={folderId || "root"}
              onUploadComplete={onUploadComplete || (() => {})}
              trigger={
                <Button variant="ghost" size="sm">
                  <Upload className="h-4 w-4 mr-2" />
                  アップロード
                </Button>
              }
            />

            <Button variant="ghost" size="sm">
              <Settings className="h-4 w-4 mr-2" />
              設定
            </Button>
            <Button variant="ghost" size="sm" onClick={onLogout}>
              <LogOut className="h-4 w-4 mr-2" />
              ログアウト
            </Button>
          </nav>
        )}
      </div>
    </header>
  )
}
