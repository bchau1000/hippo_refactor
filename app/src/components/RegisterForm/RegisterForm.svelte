<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { Events } from '../../api/constants';
	import { SimpleInput, HiddenInput, SimpleButton } from '../General/index';

	const dispatch = createEventDispatcher();

	let displayName: string = '';
	let email: string = '';
	let password: string = '';
	let confirmPassword: string = '';

	const register = () => {
		if (password !== confirmPassword) {
			console.log('Passwords do not match.');
		} else if (password.length < 8) {
			console.log('Password must be at least 8 characters.');
		} else {
			dispatch(Events.register, {
				displayName: displayName,
				email: email,
				password: password
			});
		}
	};
</script>

<form id="register-form-component" class="flex flex-col w-full gap-5">
	<div class="flex flex-col">
		<label class={`mb-2`} for="display-name">Display Name</label>
		<SimpleInput id="display-name" bind:value={displayName} />
	</div>
	<div class="flex flex-col">
		<label class={`mb-2`} for="email">Email</label>
		<SimpleInput id="email" bind:value={email} />
	</div>
	<div class="flex flex-col">
		<label class={`mb-2`} for="password">Password</label>
		<HiddenInput id="password" bind:value={password} />
	</div>
	<div class="flex flex-col mb-6">
		<label class={`mb-2`} for="confirm-password">Confirm Password</label>
		<HiddenInput id="confirm-password" bind:value={confirmPassword} />
	</div>
	<SimpleButton
		on:click={(event) => {
			event.preventDefault();
			register();
		}}
	>
		Register
	</SimpleButton>
</form>
