<script lang="ts">
	let style:string = "";
	export { style as class };
	export let value: string = '';
	export let id: string = '';
</script>

<input
	class={`${style}
        outline-none
		border-b-2
		border-gray-300
		transition-colors
		ease-in-out
		duration-200
		focus:border-yellow-400`}
	{id}
	bind:value
/>
