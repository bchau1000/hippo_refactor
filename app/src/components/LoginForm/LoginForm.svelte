<script lang="ts">
	import { Events } from '../../api/constants';
	import { createEventDispatcher } from 'svelte';
	import { SimpleInput, HiddenInput, SimpleButton } from '../General';

	let email: string = '';
	let password: string = '';

	const dispatch = createEventDispatcher();
	const onLogin = ():void => {
		dispatch(Events.login, {
			email: email,
			password: password,
		})
	}

	
</script>

<form id="login-form-component" class="flex flex-col gap-5 w-full">
	<div class="flex flex-col">
		<label class={`mb-2`} for="login-email">Email</label>
		<SimpleInput id="login-email" bind:value={email}/>
	</div>
	<div class="flex flex-col mb-6">
		<label class={`mb-2`} for="login-password">Password</label>
		<HiddenInput id="login-password" bind:value={password}/>
	</div>
	<SimpleButton
		on:click={(event) => {
			event.preventDefault();
			onLogin();
		}}
	>
		Login
	</SimpleButton>
</form>
