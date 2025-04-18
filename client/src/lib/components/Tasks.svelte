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
    Disconnect
</button>
<h1>Welcome {$user.username}</h1>
<TaskAdd />
{#if $tasks.length > 0}
    {#each $tasks as task}
        <Task {task} />
    {/each}
{:else if isFetching}
    <progress></progress>
{:else}
    <p>No tasks available.</p>
{/if}