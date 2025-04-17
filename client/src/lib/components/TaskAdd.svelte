<script lang="ts">
    import { tasks } from "$lib";
    import { fetchAPI } from "$lib/_core";

    let form:HTMLFormElement
    let title:string = $state("")
    let isSubmit:boolean = $state(false)

    async function submitTask(data:any):Promise<void>{
        data.preventDefault()
        isSubmit = true
        const response = await fetchAPI("/tasks", "POST", new FormData(form))
        isSubmit = false
        if(response.error){
            console.error(response.error)
            return
        }
        title = ""
        tasks.set(response)
    }

</script>

<form id="setTask" method="POST" onsubmit={submitTask} bind:this={form}>
    <input type="text" name="title" placeholder="Titre de la tâche" bind:value={title} />
    <button type="submit" disabled={title === "" || isSubmit === true}>Valid</button>
</form>