<script lang="ts">
	let style:string = "";
	export { style as class };
</script>

<button
	class={`${style}
        w-full 
		text-center 
		bg-yellow-400 
		py-1.5 
		rounded 
		hover:cursor-pointer 
		transition 
		ease-in-out 
		hover:bg-opacity-70 
		duration-200 `}
	on:click
>
	<slot />
</button>
