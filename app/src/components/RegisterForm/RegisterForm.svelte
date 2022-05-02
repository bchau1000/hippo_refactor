<script lang="ts">
	import { createEventDispatcher } from 'svelte';

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
			dispatch('register', {
				displayName: displayName,
				email: email,
				password: password
			});
		}
	};
</script>

<form id="register-form-component" class="flex flex-col w-full gap-5">
	<div class="flex flex-col">
		<label for="display-name">Display Name</label>
		<input
			class="outline-none 
                border-b-2 
                border-gray-300 

                transition-colors
                ease-in-out
                duration-200
                focus:border-yellow-400"
			id="display-name"
			type="name"
			bind:value={displayName}
		/>
	</div>
	<div class="flex flex-col">
		<label for="email">Email</label>
		<input
			class="outline-none 
                border-b-2 
                border-gray-300 

                transition-colors
                ease-in-out
                duration-200
                focus:border-yellow-400"
			id="email"
			type="email"
			bind:value={email}
		/>
	</div>
	<div class="flex flex-col">
		<label for="password">Password</label>
		<input
			class="outline-none 
                    border-b-2 
                    border-gray-300 
                    
                    transition-colors
                    ease-in-out
                    duration-200
                    focus:border-yellow-400"
			id="password"
			type="password"
			bind:value={password}
		/>
	</div>
	<div class="flex flex-col">
		<label for="confirm-password">Confirm Password</label>
		<input
			class="outline-none 
                    border-b-2 
                    border-gray-300 
                    
                    transition-colors
                    ease-in-out
                    duration-200
                    focus:border-yellow-400"
			id="confirm-password"
			type="password"
			bind:value={confirmPassword}
		/>
	</div>
	<input
		class="w-full 
                text-center 
                bg-yellow-400 
                mt-2
                py-1.5
                rounded 
                hover:cursor-pointer
                
                transition
                ease-in-out
                hover:bg-opacity-70
                duration-200"
		value="Register"
		type="submit"
		on:click={(event) => {
			event.preventDefault();
			register();
		}}
	/>
</form>
