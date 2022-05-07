<script lang="ts">
	import LoginForm from '../components/LoginForm';
	import RegisterForm from '../components/RegisterForm';
	import { login, register } from '../api/handleUser';

	async function handleLogin(event: any): Promise<void> {
		await login(event);
	}

	async function handleRegister(event: any):Promise<void> {
		await register(event);
	}
</script>

<main class="min-h-screen">
	<div class="flex flex-row gap-5">
		<div class="w-[500px] h-fit bg-white p-5 rounded">
			<LoginForm on:login={handleLogin}/>
		</div>
		<div class="w-[500px] h-fit bg-white p-5 rounded">
			<RegisterForm on:register={handleRegister}/>
		</div>
	</div>
</main>