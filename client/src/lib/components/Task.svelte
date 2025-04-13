<script lang="ts">
    import { tasks } from "$lib";
    import { fetchAPI } from "$lib/_core";
    import { onMount } from "svelte";

    let { task } = $props()


    async function patchTask(){
        const data = await fetchAPI(`/tasks/${task.id}`, "PATCH")
        if(data.error){
            return
        }
        task = data.result
        console.log(task, data)
    }

    async function removeTask(){
        const data = await fetchAPI(`/tasks/${task.id}`, "DELETE")
        if(data.error){
            return
        }
        tasks.update((list: { id: string; isDone: boolean; title: string }[]) => list.filter((t) => t.id !== task.id))
    }

    let dateTo:Date
    let dateToFormat:string = $state("")

    function dateFormat(){
        dateTo = new Date(task.dateTo)
        dateToFormat = dateTo.getDate()+"/"+dateTo.getMonth()+"/"+dateTo.getFullYear()+" "+dateTo.getHours()+":"+dateTo.getMinutes()
    }

    const dateAdd = new Date(task.dateAdd)
    const dateAddFormat = dateAdd.getDate()+"/"+dateAdd.getMonth()+"/"+dateAdd.getFullYear()+" "+dateAdd.getHours()+":"+dateAdd.getMinutes()


    $effect(() => {
        dateFormat()
    })

    onMount(() => {
        if(task.dateTo !== ""){
            dateFormat()
        }
    })

</script>

<tr>
    <th scope="row">{task.title}</th>
    <td>{dateAddFormat}</td>
    <td>{#if task.dateTo !== ""}{dateToFormat}{/if}</td>
    <td>
        <input type="checkbox" 
        bind:checked={task.isDone} 
        onchange={patchTask} />
    </td>
    <td>
        <button class="outline secondary" onclick={removeTask}>X</button>
    </td>
</tr>