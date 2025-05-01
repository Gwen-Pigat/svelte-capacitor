<script lang="ts">
    import { tasks } from "$lib";
    import { fetchAPI } from "$lib/_core";
    import { onMount } from "svelte";

    let { task = $bindable() } = $props()
    let isSubmit:boolean = $state(false)

    async function patchTask(){
        isSubmit = true
        const data = await fetchAPI(`/tasks/${task.id}`, "PATCH")
        isSubmit = false
        if(data.error){
            return
        }
        task = data.result
        dateFormat()
    }

    async function removeTask(){
        isSubmit = true
        const data = await fetchAPI(`/tasks/${task.id}`, "DELETE")
        isSubmit = false
        if(data.error){
            return
        }
        task.isDeleted = true
    }

    let dateTo:Date
    let dateToFormat:string = $state("")
    let diffLabel:string = $state("")

    function dateFormat(){
        if(task.dateTo === null){
            return
        }
        dateTo = new Date(task.dateTo)
        dateToFormat = dateTo.getDate()+"/"+dateTo.getMonth()+"/"+dateTo.getFullYear()+" "+dateTo.getHours()+":"+dateTo.getMinutes()+":"+dateTo.getSeconds()
        const diffMs = dateTo.getTime() - dateAdd.getTime()
        const diffSecs  = Math.floor(diffMs / 1000)
        const diffMinutes = Math.floor(diffSecs / 60)
        const hours = Math.floor(diffMinutes / 60)


        const minutes = diffMinutes % 60
        const seconds = diffSecs % 60

        console.log(seconds)

        diffLabel = "Done in "
        if(hours > 0){
            diffLabel += hours+" hour"
            if(hours > 1) diffLabel += "s"
            if(minutes > 0 || seconds > 0) diffLabel +=" et "
        }
        if(minutes > 0){
            diffLabel += minutes+" minute"
            if(minutes > 1) diffLabel += "s"
            if(seconds > 0) diffLabel += " et "
        }
        if(seconds > 0){
            diffLabel += seconds+" second"
            if(seconds > 1) diffLabel += "s"
        }
    }

    let dateAdd:any
    let dateAddFormat:any = $state("")

    $effect(() => {
        dateAdd = new Date(task.dateAdd)
        dateAddFormat = dateAdd.getDate()+"/"+dateAdd.getMonth()+"/"+dateAdd.getFullYear()+" "+dateAdd.getHours()+":"+dateAdd.getMinutes()+":"+dateAdd.getSeconds()
    })

    onMount(() => {
        dateFormat()
    })

</script>

<article>
    <header>
        {task.title} 
        {#if task.isDone}({diffLabel}){/if}
    </header>
    Add on {dateAddFormat}
    {#if task.dateTo !== null}<br />Done the {dateToFormat}{/if}
    <footer role="group">
        <button class="secondary" 
        onclick={patchTask} 
        disabled={isSubmit === true}>
            {#if task.dateTo}Cancel{:else}Finish{/if}
        </button>
        <button class="outline secondary" 
        onclick={removeTask} 
        disabled={isSubmit === true}>Delete</button>
    </footer>
</article>