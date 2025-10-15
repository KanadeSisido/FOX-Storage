"use client";

import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Folder, Home } from "lucide-react";
import Link from "next/link";

interface SidebarProps {
	currentFolderId?: string;
	folders?: Array<{ id: string; name: string; parentId?: string }>;
	onFolderCreated?: () => void;
}

export function Sidebar({
	currentFolderId,
	folders = [],
	onFolderCreated,
}: SidebarProps) {
	return (
		<aside className='w-64 border-r border-sidebar-border bg-sidebar'>
			<div className='p-4'>
				<ScrollArea className='h-[calc(100vh-8rem)]'>
					<div className='space-y-2'>
						<Button
							variant={!currentFolderId ? "secondary" : "ghost"}
							className='w-full justify-start'
							asChild
						>
							<Link href='/dashboard'>
								<Home className='h-4 w-4 mr-2' />
								ホーム
							</Link>
						</Button>

						{folders.map((folder) => (
							<Button
								key={folder.id}
								variant={currentFolderId === folder.id ? "secondary" : "ghost"}
								className='w-full justify-start'
								asChild
							>
								<Link href={`/s/${folder.id}`}>
									<Folder className='h-4 w-4 mr-2' />
									{folder.name}
								</Link>
							</Button>
						))}
					</div>
				</ScrollArea>
			</div>
		</aside>
	);
}
