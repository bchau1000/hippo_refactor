<script lang="ts">
	import '../app.css';
	import { LinkButton } from '../components/General';
</script>

<nav
	class={`bg-dark-fg 
        absolute
        w-[100vw]
        max-w-full
        top-0
        h-[60px]
        flex
        flex-row
        justify-between`}
>
	<div
        class={`pl-6`}
    >
		<a
			href="/"
			class={`h-full 
            flex 
            flex-row 
            items-center
            text-light-fg
            hover:text-yellow-400
            duration-200`}
		>
			<span
				class={`h-fit
                font-bold
                text-3xl
            `}
			>
				Hippo.
			</span>
		</a>
	</div>

	<div class={`flex flex-row`}>
		<LinkButton href="/version">Browse</LinkButton>
        <LinkButton href="/version">Sets</LinkButton>
        <LinkButton href="/version">Create</LinkButton>
	</div>
</nav>

<main 
    class={`w-[100vw] 
        max-w-full 
        h-[100vh]
        bg-slate-200
        pt-[90px]`}
    >
    <slot />
</main>