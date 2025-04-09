<script lang="ts">
    import Task from '$lib/components/Task.svelte';
    import TaskAdd from '$lib/components/TaskAdd.svelte';
    import { tasks } from '$lib';

    import { onMount } from 'svelte';

    let user:any

    async function loadUser(){
        const data = await fetch("/api/auth")
        const result = await data.json()
        if(!data.ok){
            return
        }
        if(result.user){
            user = result.user
            loadData()
        }
    }

    async function loadData(){
        const data = await fetch("/api/tasks")
        tasks.set(await data.json())
        console.log($tasks.length)
    }

    async function connectUser(data:any){
        data.preventDefault()
        const dataFetch = await fetch("/api/auth", {
            method: "POST",
            body: new FormData(data.currentTarget)
        })

        const result = await dataFetch.json()
        if(!dataFetch.ok){
            console.error(result.error)
            return
        }
        user = result.user
        console.log("Response cookies", result, user)
    }

    onMount(() => {
        loadUser()
    })

</script>


{#if !user}
    <h1>Connect</h1>
    <form id="setConnect" onsubmit={connectUser}>
        <input type="text" placeholder="Votre nom" name="user" />
        <button type="submit">Connect</button>
    </form>
{:else}
    <h1>Tasks Loader</h1>
    <TaskAdd />
    {#if $tasks.length > 0}
        <table>
            <thead>
                <tr>
                    <th scope="col">Title</th>
                    <th scope="col">Date add</th>
                    <th scope="col">Date done</th>
                    <th scope="col">Is done</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                {#each $tasks as task}
                    <Task {task} />
                {/each}
            </tbody>
        </table>
    {:else}
        <p>No tasks available.</p>
    {/if}
{/if}