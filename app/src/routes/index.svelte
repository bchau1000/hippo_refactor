<script lang="ts">
	import LoginForm from '../components/LoginForm/LoginForm.svelte';
	import { login } from '../api/handleUser';

	async function handleLogin(event: any): Promise<void> {
		const response:Response = await login(event);
		
		if(response.status === 200) {
			const json:JSON = await response.json();
			console.log(json);
		}
	}
</script>

<main class="bg-slate-200 min-h-screen">
	<div class="w-[500px] bg-white p-5 rounded">
		<LoginForm on:message={handleLogin}/>
	</div>
</main>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Fira+Sans:wght@400;700&display=swap');
	main {
		font-family: 'Fira Sans', sans-serif !important;
	}
</style>
