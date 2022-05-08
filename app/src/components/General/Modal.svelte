<script lang="ts">
    import { fly, fade } from 'svelte/transition';

    let style:string = "";
	export { style as class };
    export let showModal:boolean = false;
    export let showFocus:boolean = false;
</script>

{#if showModal}
{#if showFocus}
<div
    class={`absolute
    w-full
    h-full
    top-0
    left-0
    bg-gray-900
    opacity-50`}
/>
{/if}
<div
    class={`${style}
        absolute
        w-fit
        h-fit
        
        mx-auto
        left-0
        right-0
        flex
        flex-col
    `}
    transition:fly="{{ y: -200, duration: 250 }}"
>
    
    <slot />
</div>
{/if}