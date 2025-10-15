"use client";
import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Header } from "@/components/layout/header";
import { Sidebar } from "@/components/layout/sidebar";
import { FileGrid } from "@/components/files/file-grid";
import { Breadcrumb } from "@/components/files/breadcrumb";
import { UploadDialog } from "@/components/files/upload-dialog";
import { CreateFolderDialog } from "@/components/files/create-folder-dialog";
import { Button } from "@/components/ui/button";
import { Upload, FolderPlus } from "lucide-react";
import { FileItem } from "@/app/dashboard/page";

interface BreadcrumbItem {
	name: string;
	path: string;
}

export default function FolderPage() {
	const params = useParams();
	const router = useRouter();
	const folderId = params.folderId as string;

	const [files, setFiles] = useState<FileItem[]>([]);
	const [folders, setFolders] = useState<
		Array<{ id: string; name: string; parentId?: string }>
	>([]);
	const [breadcrumbs, setBreadcrumbs] = useState<BreadcrumbItem[]>([]);
	const [isLoading, setIsLoading] = useState(true);
	const [currentFolder, setCurrentFolder] = useState<{
		id: string;
		name: string;
	} | null>(null);

	useEffect(() => {
		if (folderId) {
			loadFolderContents(folderId);
		}
	}, [folderId]);

	const loadFolderContents = async (id: string) => {
		setIsLoading(true);
		try {
			const response = await fetch(`http://localhost:8080/s/${folderId}/`, {
				method: "GET",
				credentials: "include",
			});
			if (response.ok) {
				const json: { items: FileItem[] } = await response.json();
				const data = json.items;

				const files = data.filter((item) => {
					return item.type === "file";
				});

				const folders = data.filter((item) => {
					return item.type === "folder";
				});

				setFiles(files || []);
				setFolders(folders || []);
			}
		} catch (error) {
			console.error("Failed to load files:", error);
		} finally {
			setIsLoading(false);
		}
	};

	const handleLogout = () => {
		localStorage.removeItem("user");
		router.push("/auth/login");
	};

	const handleFileClick = (file: FileItem) => {
		if (file.type === "folder") {
			router.push(`/s/${file.id}`);
		}
	};

	const handleDownload = async (fileId: string, fileName: string) => {
		try {
			const response = await fetch(
				`http://localhost:8080/s/${folderId}/file/${fileId}`,
				{
					method: "GET",
					credentials: "include",
				}
			);
			if (response.ok) {
				const blob = await response.blob();
				const url = window.URL.createObjectURL(blob);
				const a = document.createElement("a");
				a.href = url;
				a.download = fileName;
				document.body.appendChild(a);
				a.click();
				window.URL.revokeObjectURL(url);
				document.body.removeChild(a);
			}
		} catch (error) {
			console.error("Download failed:", error);
		}
	};

	const refreshFiles = () => {
		if (folderId) {
			loadFolderContents(folderId);
		}
	};

	return (
		<AuthGuard>
			<div className='min-h-screen bg-background'>
				<Header
					showNavigation
					onLogout={handleLogout}
					onUploadComplete={refreshFiles}
				/>
				<div className='flex'>
					<Sidebar
						currentFolderId={folderId}
						folders={folders}
						onFolderCreated={refreshFiles}
					/>
					<main className='flex-1 p-6'>
						<div className='max-w-7xl mx-auto'>
							<Breadcrumb items={breadcrumbs} />

							<div className='mt-6'>
								<div className='flex items-center justify-between mb-6'>
									<h2 className='text-2xl font-bold text-foreground'>
										{currentFolder?.name || "フォルダ"}
									</h2>

									<div className='flex gap-2'>
										<CreateFolderDialog
											parentFolderId={folderId}
											onFolderCreated={refreshFiles}
											trigger={
												<Button variant='outline'>
													<FolderPlus className='h-4 w-4 mr-2' />
													新しいフォルダ
												</Button>
											}
										/>
										<UploadDialog
											folderId={folderId}
											onUploadComplete={refreshFiles}
											trigger={
												<Button>
													<Upload className='h-4 w-4 mr-2' />
													ファイルをアップロード
												</Button>
											}
										/>
									</div>
								</div>

								{isLoading ? (
									<div className='flex items-center justify-center py-12'>
										<div className='text-center'>
											<div className='animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4'></div>
											<p className='text-muted-foreground'>
												ファイルを読み込み中...
											</p>
										</div>
									</div>
								) : (
									<FileGrid
										files={files}
										onFileClick={handleFileClick}
										onDownload={handleDownload}
									/>
								)}
							</div>
						</div>
					</main>
				</div>
			</div>
		</AuthGuard>
	);
}
