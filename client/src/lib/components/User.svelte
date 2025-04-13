<script lang="ts">
    import { user } from "$lib";
    import { fetchAPI } from "$lib/_core";
    import { Preferences } from '@capacitor/preferences';

    let formView = $state("default")

    let form:any = $state()

    async function fetchUser(data:any){
        data.preventDefault()
        const dataFetch = await fetchAPI(
            form.getAttribute("data-path"), 
            "POST", 
            new FormData(form)
        )
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
    class="action outline" 
    onclick={() => formView = "register"}>
        Register
    </button>
    <form id="setConnect" 
    bind:this={form}
    data-path="/user/connect" 
    onsubmit={fetchUser}>
        <input type="text" placeholder="Votre nom" name="username" />
        <button type="submit">Connect</button>
    </form>
{:else if formView === "register"}
    <h1>Register</h1>
    <button type="button" 
    class="action outline" 
    onclick={() => formView = "default"}>
        Connect
    </button>
    <form id="setConnect" 
    bind:this={form}
    data-path="/user" 
    onsubmit={fetchUser}>
        <input type="text" placeholder="Votre nom" name="username" />
        <button type="submit">Register</button>
    </form>
{/if}