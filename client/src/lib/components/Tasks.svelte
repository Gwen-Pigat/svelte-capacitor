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
class="action outline" 
onclick={resetUser}>
    Disconnect
</button>
<h1>Bienvenue {$user.username}</h1>
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