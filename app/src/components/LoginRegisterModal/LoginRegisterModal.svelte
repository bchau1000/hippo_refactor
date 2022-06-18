<script lang="ts">
	import { Modal } from '../General';
	import LoginForm from '../LoginForm';
	import RegisterForm from '../RegisterForm';

	export let showModal: boolean = false;
	let showTab: boolean = false;

</script>

<Modal class={`flex flex-col `} bind:showModal showFocus={true}>
	<section class={`h-fit !bg-transparent flex flex-row-reverse`}>
		<button class={`text-slate-200 py-1 pl-2`} on:click={() => (showModal = false)}>
			<span class="material-symbols-outlined text-[38px] hover:text-yellow-400">close</span>
		</button>
	</section>
	<section class={`flex w-full bg-white`}>
		<button
			class={`w-[50%] p-4
                ${showTab ? '' : 'bg-gray-100'}`}
			on:click={() => (showTab = true)}
		>
			Login
		</button>
		<button
			class={`w-[50%] p-4
                ${showTab ? 'bg-gray-100' : ''}`}
			on:click={() => (showTab = false)}
		>
			Register
		</button>
	</section>
	<section
		class={`bg-white
        w-[500px]
        px-12
        pt-12
        pb-8O
    `}
	>
		{#if showTab}
			<LoginForm on:login/>
		{:else}
			<RegisterForm on:register/>
		{/if}
	</section>
</Modal>
