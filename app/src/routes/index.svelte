<script lang="ts">
	import {Version} from '../models/Version';
	import { ServerInfo, ContentType } from "../api/constants";
	async function request():Promise<Version> {
		const response:Response = await fetch(
			"http://127.0.0.1:4000/api/version", 
			{
				method: "GET",
				credentials: 'include',
				headers: {
                    ...ContentType.json,
                },
			});
		const json = await response.json();
		return new Version(json);
	}
	let version:Promise<Version> = request();
</script>

<main class="min-h-[100%]">
	{#await version}
		<span class="text-med">Loading...</span>
	{:then version} 
		<span class="text-med">{version.toString()}</span>
	{/await}
</main>