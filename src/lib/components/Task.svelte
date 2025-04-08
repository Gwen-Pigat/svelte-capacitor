<script lang="ts">
    import { tasks } from "$lib";
    import { onMount } from "svelte";

    let { task } = $props()


    async function loadData(){
        const data = await fetch(`/api/tasks/${task.id}`, {
            method: "PATCH"
        })
        const response = await data.json()
        if(response.error){
            console.error(response.error)
        }
        task = response.task
    }

    async function removeTask(){
        const data = await fetch(`/api/tasks/${task.id}`, {
            method: "DELETE"
        })
        const response = await data.json()
        if(response.error){
            console.error(response.error)
        }
        tasks.update((list: { id: string; is_done: boolean; title: string }[]) => list.filter((t) => t.id !== task.id))
    }

    let dateTo:Date
    let dateToFormat:string = $state("")

    function dateFormat(date:Date){
        dateTo = new Date(task.date_to)
        dateToFormat = dateTo.getDate()+"/"+dateTo.getMonth()+"/"+dateTo.getFullYear()+" "+dateTo.getHours()+":"+dateTo.getMinutes()
    }

    const dateAdd = new Date(task.date_add)
    const dateAddFormat = dateAdd.getDate()+"/"+dateAdd.getMonth()+"/"+dateAdd.getFullYear()+" "+dateAdd.getHours()+":"+dateAdd.getMinutes()


    $effect(() => {
        dateFormat(task.date_to)
    })

    onMount(() => {
        if(task.date_to !== null){
            dateFormat(task.date_to)
        }
    })

</script>

<tr>
    <th scope="row">{task.title}</th>
    <td>{dateAddFormat}</td>
    <td>{#if task.date_to !== null}{dateToFormat}{/if}</td>
    <td>
        <input type="checkbox" 
        bind:checked={task.is_done} 
        onchange={loadData} />
    </td>
    <td>
        <button class="outline secondary" onclick={removeTask}>X</button>
    </td>
</tr>