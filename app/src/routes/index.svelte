<script lang="ts">
	import {Version} from '../models/Version';
	async function request():Promise<Version> {
		const response:Response = await fetch("http://localhost:3000/api/version");
		const json = await response.json();
		return new Version(json);
	}
	let version:Promise<Version> = request();
</script>

<main class="min-h-[100%] bg-dark-bg ">
	{#await version}
		<span class="text-med">Loading...</span>
	{:then version} 
		<span class="text-med">{version.toString()}</span>
	{/await}
</main>