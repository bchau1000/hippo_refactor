<script lang="ts">
    let style:string = "";
	export { style as class };
    export let href = "/";
</script>

<a 
    href={href}
    class={`${style}
        h-full
        font-bold
        flex
        flex-row
        items-center
        text-light-fg
        hover:text-dark-fg
        hover:bg-yellow-400
        duration-200
        px-6
    `}
>
    <span
        class={`h-fit
            text-xl
        `}
    ><slot/></span>
</a>