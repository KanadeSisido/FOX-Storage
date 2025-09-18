"use client";

import type React from "react";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
	Dialog,
	DialogContent,
	DialogDescription,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { FolderPlus } from "lucide-react";

interface CreateFolderDialogProps {
	parentFolderId?: string;
	onFolderCreated: () => void;
	trigger?: React.ReactNode;
}

export function CreateFolderDialog({
	parentFolderId,
	onFolderCreated,
	trigger,
}: CreateFolderDialogProps) {
	const [isOpen, setIsOpen] = useState(false);
	const [folderName, setFolderName] = useState("");
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState("");
	console.log(parentFolderId);
	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		if (!folderName.trim()) {
			setError("フォルダ名を入力してください");
			return;
		}

		setIsLoading(true);
		setError("");

		try {
			const response = await fetch(
				`http://localhost:8080/s/${parentFolderId}/folder/create`,
				{
					method: "POST",
					headers: {
						"Content-Type": "application/json",
					},
					body: JSON.stringify({
						name: folderName.trim(),
					}),
					credentials: "include",
				}
			);

			if (response.ok) {
				onFolderCreated();
				setIsOpen(false);
				setFolderName("");
			} else {
				const errorData = await response.json();
				setError(errorData.message || "フォルダの作成に失敗しました");
			}
		} catch (err) {
			setError("ネットワークエラーが発生しました");
		} finally {
			setIsLoading(false);
		}
	};

	return (
		<Dialog open={isOpen} onOpenChange={setIsOpen}>
			<DialogTrigger asChild>
				{trigger || (
					<Button variant='outline'>
						<FolderPlus className='h-4 w-4 mr-2' />
						新しいフォルダ
					</Button>
				)}
			</DialogTrigger>
			<DialogContent className='max-w-md'>
				<DialogHeader>
					<DialogTitle>新しいフォルダを作成</DialogTitle>
					<DialogDescription>フォルダ名を入力してください</DialogDescription>
				</DialogHeader>

				<form onSubmit={handleSubmit} className='space-y-4'>
					{error && (
						<Alert variant='destructive'>
							<AlertDescription>{error}</AlertDescription>
						</Alert>
					)}

					<div className='space-y-2'>
						<Label htmlFor='folderName'>フォルダ名</Label>
						<Input
							id='folderName'
							type='text'
							placeholder='新しいフォルダ'
							value={folderName}
							onChange={(e) => setFolderName(e.target.value)}
							required
						/>
					</div>

					<div className='flex justify-end gap-2'>
						<Button
							type='button'
							variant='outline'
							onClick={() => {
								setIsOpen(false);
								setFolderName("");
								setError("");
							}}
							disabled={isLoading}
						>
							キャンセル
						</Button>
						<Button type='submit' disabled={isLoading}>
							{isLoading ? "作成中..." : "作成"}
						</Button>
					</div>
				</form>
			</DialogContent>
		</Dialog>
	);
}
