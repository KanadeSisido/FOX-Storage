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

export default function SignupPage() {
	const [username, setUsername] = useState("");
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [confirmPassword, setConfirmPassword] = useState("");
	const [showPassword, setShowPassword] = useState(false);
	const [showConfirmPassword, setShowConfirmPassword] = useState(false);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState("");
	const router = useRouter();

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		setIsLoading(true);
		setError("");

		if (password !== confirmPassword) {
			setError("パスワードが一致しません");
			setIsLoading(false);
			return;
		}

		if (password.length < 8) {
			setError("パスワードは8文字以上で入力してください");
			setIsLoading(false);
			return;
		}

		try {
			const response = await fetch("http://localhost:8080/auth/signup", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ username, email, password }),
			});

			if (response.ok) {
				await response.json();
				router.push("/auth/login");
			} else {
				const errorData = await response.json();
				setError(errorData.message || "アカウント作成に失敗しました");
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
						<CardTitle className='text-2xl font-bold'>新規登録</CardTitle>
						<CardDescription>
							新しいアカウントを作成してファイルストレージを始めましょう
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
								<Label htmlFor='name'>ユーザ名</Label>
								<Input
									id='name'
									type='text'
									placeholder='username'
									value={username}
									onChange={(e) => setUsername(e.target.value)}
									required
								/>
							</div>

							<div className='space-y-2'>
								<Label htmlFor='email'>メールアドレス</Label>
								<Input
									id='email'
									type='email'
									placeholder='your@email.com'
									value={email}
									onChange={(e) => setEmail(e.target.value)}
									required
								/>
							</div>

							<div className='space-y-2'>
								<Label htmlFor='password'>パスワード</Label>
								<div className='relative'>
									<Input
										id='password'
										type={showPassword ? "text" : "password"}
										placeholder='8文字以上のパスワード'
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

							<div className='space-y-2'>
								<Label htmlFor='confirmPassword'>パスワード確認</Label>
								<div className='relative'>
									<Input
										id='confirmPassword'
										type={showConfirmPassword ? "text" : "password"}
										placeholder='パスワードを再入力'
										value={confirmPassword}
										onChange={(e) => setConfirmPassword(e.target.value)}
										required
									/>
									<Button
										type='button'
										variant='ghost'
										size='sm'
										className='absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent'
										onClick={() => setShowConfirmPassword(!showConfirmPassword)}
									>
										{showConfirmPassword ? (
											<EyeOff className='h-4 w-4' />
										) : (
											<Eye className='h-4 w-4' />
										)}
									</Button>
								</div>
							</div>

							<Button type='submit' className='w-full' disabled={isLoading}>
								{isLoading ? "アカウント作成中..." : "アカウント作成"}
							</Button>
						</form>

						<div className='mt-6 text-center text-sm'>
							<span className='text-muted-foreground'>
								すでにアカウントをお持ちの方は{" "}
							</span>
							<Link href='/auth/login' className='text-primary hover:underline'>
								ログイン
							</Link>
						</div>
					</CardContent>
				</Card>
			</div>
		</div>
	);
}
