"use client";

import type React from "react";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from "@/components/ui/card";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Header } from "@/components/layout/header";
import { Cloud, Eye, EyeOff } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";

export default function LoginPage() {
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [showPassword, setShowPassword] = useState(false);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState("");
	const router = useRouter();

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		setIsLoading(true);
		setError("");

		try {
			const response = await fetch("http://localhost:8080/auth/login", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ username, password }),
				credentials: "include",
			});

			if (response.ok) {
				const data = await response.json();
				console.log(response);
				// Store user data and redirect to dashboard
				localStorage.setItem("user", JSON.stringify(username));
				router.push("/dashboard");
			} else {
				const errorData = await response.json();
				setError(errorData.message || "ログインに失敗しました");
			}
		} catch (err) {
			setError("ネットワークエラーが発生しました");
		} finally {
			setIsLoading(false);
		}
	};

	return (
		<div className='min-h-screen bg-background'>
			<Header />
			<div className='flex items-center justify-center min-h-[calc(100vh-4rem)] p-4'>
				<Card className='w-full max-w-md'>
					<CardHeader className='text-center'>
						<div className='flex justify-center mb-4'>
							<Cloud className='h-12 w-12 text-primary' />
						</div>
						<CardTitle className='text-2xl font-bold'>ログイン</CardTitle>
						<CardDescription>
							アカウントにログインしてファイルにアクセスしましょう
						</CardDescription>
					</CardHeader>
					<CardContent>
						<form onSubmit={handleSubmit} className='space-y-4'>
							{error && (
								<Alert variant='destructive'>
									<AlertDescription>{error}</AlertDescription>
								</Alert>
							)}

							<div className='space-y-2'>
								<Label htmlFor='username'>ユーザ名</Label>
								<Input
									id='username'
									type='username'
									placeholder='ユーザ名'
									value={username}
									onChange={(e) => setUsername(e.target.value)}
									required
								/>
							</div>

							<div className='space-y-2'>
								<Label htmlFor='password'>パスワード</Label>
								<div className='relative'>
									<Input
										id='password'
										type={showPassword ? "text" : "password"}
										placeholder='パスワードを入力'
										value={password}
										onChange={(e) => setPassword(e.target.value)}
										required
									/>
									<Button
										type='button'
										variant='ghost'
										size='sm'
										className='absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent'
										onClick={() => setShowPassword(!showPassword)}
									>
										{showPassword ? (
											<EyeOff className='h-4 w-4' />
										) : (
											<Eye className='h-4 w-4' />
										)}
									</Button>
								</div>
							</div>

							<Button type='submit' className='w-full' disabled={isLoading}>
								{isLoading ? "ログイン中..." : "ログイン"}
							</Button>
						</form>

						<div className='mt-6 text-center text-sm'>
							<span className='text-muted-foreground'>
								アカウントをお持ちでない方は{" "}
							</span>
							<Link
								href='/auth/signup'
								className='text-primary hover:underline'
							>
								新規登録
							</Link>
						</div>
					</CardContent>
				</Card>
			</div>
		</div>
	);
}
