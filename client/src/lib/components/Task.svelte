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
    let diffLabel:string = $state("")

    function dateFormat(){
        dateTo = new Date(task.dateTo)
        dateToFormat = dateTo.getDate()+"/"+dateTo.getMonth()+"/"+dateTo.getFullYear()+" "+dateTo.getHours()+":"+dateTo.getMinutes()
        const diffMs = dateTo.getTime() - dateAdd.getTime()
        const diffMins  = Math.floor(diffMs / 60000)
        const hours = Math.floor(diffMins / 60)
        const minutes = diffMins % 60

        diffLabel = "Finie en "
        if(hours > 0){
            diffLabel += hours+" heure"
            if(hours > 1) diffLabel += "s"
            if(minutes > 0) diffLabel +=" et "
        }
        if(minutes > 0){
            diffLabel += minutes+" minute"
            if(minutes > 1) diffLabel += "s"
        }

    }

    const dateAdd = new Date(task.dateAdd)
    const dateAddFormat = dateAdd.getDate()+"/"+dateAdd.getMonth()+"/"+dateAdd.getFullYear()+" "+dateAdd.getHours()+":"+dateAdd.getMinutes()

    onMount(() => {
        if(task.dateTo !== ""){
            dateFormat()
        }
    })

</script>

<article>
    <header>
        {task.title} 
        {#if task.isDone}{diffLabel}{/if}
    </header>
    Ajoutée le {dateAddFormat}
    {#if task.dateTo !== ""}<br />Finie le {dateToFormat}{/if}
    <footer>
        <button class="secondary" onclick={patchTask}>
            {#if task.dateTo}Annuler{:else}Finir{/if}
        </button>
        <button class="outline secondary" onclick={removeTask}>Supprimer</button>
    </footer>
</article>