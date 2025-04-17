<script lang="ts">
    import { user } from "$lib";
    import { fetchAPI } from "$lib/_core";
    import { Preferences } from '@capacitor/preferences';

    let formView = $state("default")

    let form:any = $state()
    let isSubmit:boolean = $state(false)

    async function fetchUser(data:any){
        data.preventDefault()
        isSubmit = true
        const dataFetch = await fetchAPI(
            form.getAttribute("data-path"), 
            "POST", 
            new FormData(form)
        )
        isSubmit = false
        if(dataFetch.error){
            return
        }
        await Preferences.set({
            key: "user",
            value: dataFetch.token
        })
        user.set(dataFetch)
    }

</script>

{#if formView === "default"}
    <h1>Connect</h1>
    <button type="button" 
    class="action outline secondary" 
    onclick={() => formView = "register"}>
        Register
    </button>
    <form id="setConnect" 
    bind:this={form}
    data-path="/user/connect" 
    onsubmit={fetchUser}>
        <input type="text" placeholder="Your name" name="username" />
        <button type="submit" disabled={isSubmit === true}>Connect</button>
    </form>
{:else if formView === "register"}
    <h1>Register</h1>
    <button type="button" 
    class="action outline secondary" 
    onclick={() => formView = "default"}>
        Connect
    </button>
    <form id="setConnect" 
    bind:this={form}
    data-path="/user" 
    onsubmit={fetchUser}>
        <input type="text" placeholder="Your name" name="username" />
        <button type="submit" disabled={isSubmit === true}>Register</button>
    </form>
{/if}