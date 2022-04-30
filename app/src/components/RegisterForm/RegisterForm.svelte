<script lang="ts">
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    let email: string = '';
	let password: string = '';

    

</script>

<form id="register-form-component" class="block w-full">
    <div class="flex flex-col pb-2">
		<label for="email">Email</label>
		<input
			class="outline-none 
                border-b-2 
                border-gray-300 
                pb-1

                transition-colors
                ease-in-out
                duration-200
                focus:border-yellow-400"
			id="email"
			type="email"
			bind:value={email}
		/>
        <div class="flex flex-col pb-5">
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
        <input
            class="w-full 
                text-center 
                bg-yellow-400 
                py-1.5
                rounded 
                hover:cursor-pointer
                
                transition
                ease-in-out
                hover:bg-opacity-70
                duration-200"
            value="Login"
            type="submit"
            on:click={(event) => {
                event.preventDefault();
                
            }}
        />
	</div>
</form>