<script lang="ts">
	import '../app.css';
    import { login, register } from '../api/handleUser';
	import { LinkButton } from '../components/General';
    import LoginRegisterModal from '../components/LoginRegisterModal/';

    let showModal:boolean = false;

    const onLogin = async (
        {email, password}: 
        {email:string, password:string}
    ):Promise<void> => {
        await login(email, password);
    }

    const onRegister = async (
        {displayName, email, password, confirmPassword}:
        {displayName:string, email:string, password:string, confirmPassword:string}
    ):Promise<void> => {
        if(password !== confirmPassword) {
            // TODO: alert/notification system
            console.log("Passwords must match");
            return;
        }

        if(password.length < 8) {
            // TODO: alert/notification system
            console.log("Password must be at least 8 characters long.");
            return;
        }

        await register(displayName, email, password);
    }

</script>

<nav class={`bg-dark-fg 
    absolute
    w-[100vw]
    max-w-full
    top-0
    h-[60px]
    flex
    flex-row
    justify-between`}
>
	<div class={`pl-6`}>
		<a href="/"
			class={`h-full 
            flex 
            flex-row 
            items-center
            text-light-fg
            hover:text-yellow-400
            duration-200`}
		>
			<span class={`h-fit
                font-bold
                text-3xl`}
			>Hippo.</span>
		</a>
	</div>

	<div class={`flex flex-row items-center`}>
		<LinkButton href="/sandbox">Browse</LinkButton>
        <LinkButton href="/sandbox">Sets</LinkButton>
        <LinkButton href="/sandbox">Create</LinkButton>
        <div class={`w-[0.2px] h-[80%] bg-gray-600 mx-4`}/>
        <button class={`h-full
            font-bold
            flex
            flex-row
            items-center
            text-light-fg
            hover:text-yellow-400
            duration-200
            text-xl
            px-6`}
            on:click={() => showModal = true}
        >
            Login/Register
        </button>
	</div>
</nav>

<main 
    class={`w-[100vw] 
        max-w-full 
        h-[100vh]
        bg-slate-200
        pt-[90px]`}
    >
    <LoginRegisterModal 
        bind:showModal={showModal}
        on:login={(event) => onLogin(event.detail)}
        on:register={(event) => onRegister(event.detail)}
    />
    <slot />
</main>