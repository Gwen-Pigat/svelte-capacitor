<script lang="ts">
    import { onMount } from 'svelte';
    import { fetchAPI } from '$lib/_core';
    import { user,error } from '$lib'
    import Tasks from '$lib/components/Tasks.svelte';
    import User from '$lib/components/User.svelte';
    import { fade } from 'svelte/transition';

    async function loadUser(){
        const data = await fetchAPI("/user","GET")
        userConnect = true
        if(data.error){
            return
        }
        user.set(data)
    }

    let userConnect:boolean = $state(false)

    onMount(() => {
        loadUser()
    })

    $effect(() => {
        if($error !== ""){
            setTimeout(function(){
                error.set("")
            },6000)
        }
    })

</script>

<img class="logo" src="/images/logo.svg" alt="Tasker Logo" />
{#if !userConnect}
    Chargement en cours
{:else}
    {#if $error !== ""}
        <div class="error" transition:fade={{duration:500}}>{$error}</div>
    {/if}
    {#if !$user.id}
        <User />
    {:else}
        <Tasks />
    {/if}
{/if}

<style>
    .error{
        position: absolute;
        right: 0;
        top: 0;
        font-size: 16px;
        padding: 0.5rem;
        border-radius: 0;
        color: var(--white);
        background-color: var(--red);
        z-index: 10000;
    }
</style>