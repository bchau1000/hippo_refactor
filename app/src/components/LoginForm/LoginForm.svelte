<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { Events } from '../../api/constants';
	import { SimpleInput, HiddenInput, SimpleButton } from '../General';

	const dispatch = createEventDispatcher();

	let email: string = '';
	let password: string = '';

	const login = () => {
		if (email.length > 0 && password.length > 0) {
			dispatch(Events.login, {
				email: email,
				password: password,
			});
		} else {
			// TODO: Alert system that isn't annoying
			console.log("Invalid email and/or password...");
		}
	}
</script>

<form id="login-form-component" class="flex flex-col gap-5 w-full">
	<div class="flex flex-col">
		<label for="email">Email</label>
		<SimpleInput id="email" bind:value={email}/>
	</div>
	<div class="flex flex-col">
		<label for="password">Password</label>
		<HiddenInput id="password" bind:value={password}/>
	</div>
	<SimpleButton
		on:click={(event) => {
			event.preventDefault();
			login();
		}}
	>
		Login
	</SimpleButton>
</form>
