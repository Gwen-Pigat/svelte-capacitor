<script lang="ts">
    import Task from '$lib/components/Task.svelte';
    import TaskAdd from '$lib/components/TaskAdd.svelte';
    import { tasks } from '$lib';

    import { onMount } from 'svelte';

    async function loadData(){
        const data = await fetch("/api/tasks")
        tasks.set(await data.json())
        console.log($tasks.length)
    }

    onMount(() => {
        loadData()
    })

    $effect(() => {
        console.log($tasks)
    })

    let addTask = $state(false)

</script>


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