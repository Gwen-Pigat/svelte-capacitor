<script lang="ts">
    import { tasks } from "$lib";
    import { fetchAPI } from "$lib/_core";

    let form:HTMLFormElement
    let title:string = ""

    async function submitTask(data:any):Promise<void>{
        data.preventDefault()
        const response = await fetchAPI("/tasks", "POST", new FormData(form))
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
    <button type="submit" disabled={title === ""}>Valid</button>
</form>