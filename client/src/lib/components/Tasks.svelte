<script lang="ts">
    import { tasks,user } from "$lib"
    import { onMount } from "svelte";
    import Task from "./Task.svelte";
    import TaskAdd from "./TaskAdd.svelte";
    import { fetchAPI, resetUser } from "$lib/_core";


    let isFetching:boolean = false

    async function loadTasks(){
        isFetching = true
        const dataFetch = await fetchAPI("/tasks", "GET")
        isFetching = false
        if(dataFetch.error){
            return
        }
        tasks.set(dataFetch)
    }

    onMount(() => {
        loadTasks()
    })

</script>

<button type="button" 
class="action outline secondary" 
onclick={resetUser}>
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M288 32c0-17.7-14.3-32-32-32s-32 14.3-32 32l0 224c0 17.7 14.3 32 32 32s32-14.3 32-32l0-224zM143.5 120.6c13.6-11.3 15.4-31.5 4.1-45.1s-31.5-15.4-45.1-4.1C49.7 115.4 16 181.8 16 256c0 132.5 107.5 240 240 240s240-107.5 240-240c0-74.2-33.8-140.6-86.6-184.6c-13.6-11.3-33.8-9.4-45.1 4.1s-9.4 33.8 4.1 45.1c38.9 32.3 63.5 81 63.5 135.4c0 97.2-78.8 176-176 176s-176-78.8-176-176c0-54.4 24.7-103.1 63.5-135.4z"/></svg> Disconnect
</button>
<h1>Welcome {$user.username}</h1>
<TaskAdd />
{#if $tasks.length > 0}
    {#each $tasks as task}
        {#if !task.isDeleted}
            <Task bind:task />
        {/if}
    {/each}
{:else if isFetching}
    <progress></progress>
{:else}
    <p>No tasks available.</p>
{/if}