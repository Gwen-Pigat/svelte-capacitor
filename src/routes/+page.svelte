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
{#if addTask}
    <TaskAdd bind:addTask />
{:else}
    <button type="button" onclick={() => addTask = !addTask}>Add task</button>
{/if}
{#if $tasks.length > 0}
    <div id="tasks">
        {#each $tasks as task}
            <Task {task} />
        {/each}
    </div>
{:else}
    <p>No tasks available.</p>
{/if}