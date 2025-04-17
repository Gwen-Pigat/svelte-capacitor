<script>
    import { tasks,user } from "$lib"
    import { onMount } from "svelte";
    import Task from "./Task.svelte";
    import TaskAdd from "./TaskAdd.svelte";
    import { fetchAPI, resetUser } from "$lib/_core";


    async function loadTasks(){
        const dataFetch = await fetchAPI("/tasks", "GET")
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
{:else}
    <p>No tasks available.</p>
{/if}