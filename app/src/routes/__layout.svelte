<script lang="ts">
	import '../app.css';
	import { afterUpdate, onMount } from 'svelte';
	import { LinkButton } from '../components/General';
	import LoginRegisterModal from '../components/LoginRegisterModal/';
	import { authenticate } from '../api/handleUser';

	let showModal: boolean = false;
	let loggedIn: boolean = false;

	onMount(():void => {
		const user:string = authenticate();
		console.log(user);
		if(user === "")
			loggedIn = false;
		else
			loggedIn = true;
	});
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
    justify-between
	shadow-md`}
>
	<div class={`pl-6`}>
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
                text-3xl`}>Hippo.</span
			>
		</a>
	</div>

	<div class={`flex flex-row items-center`}>
		<LinkButton href="/sandbox">Browse</LinkButton>
		<LinkButton href="/sandbox">Sets</LinkButton>
		<LinkButton href="/sandbox">Create</LinkButton>
		<div class={`w-[0.2px] h-[80%] bg-gray-600 mx-4`} />
		
		{#if loggedIn}
			<button
				class={`h-full
        	    font-bold
        	    flex
        	    flex-row
        	    items-center
        	    text-light-fg
        	    hover:text-yellow-400
        	    duration-200
        	    text-xl
        	    px-6`}
			>
				Logout
			</button>
		{:else}
			<button
				class={`h-full
        	    font-bold
        	    flex
        	    flex-row
        	    items-center
        	    text-light-fg
        	    hover:text-yellow-400
        	    duration-200
        	    text-xl
        	    px-6`}
				on:click={() => (showModal = true)}
			>
				Login/Register
			</button>
		{/if}
	</div>
</nav>

<main
	class={`w-[100vw] 
        max-w-full 
        h-[100vh]
        bg-slate-200`}
>
	<LoginRegisterModal bind:showModal />
	<slot />
</main>
