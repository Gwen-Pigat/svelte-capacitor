<script lang="ts">
    let { task } = $props()


    async function loadData(){
        const data = await fetch(`/api/tasks/${task.id}`, {
            method: "PATCH"
        })
        const response = await data.json()
        if(response.error){
            console.error(response.error)
        }
    }

</script>

<div class="task">
    <h2>{task.title}</h2>
    <input type="checkbox" 
    bind:checked={task.is_done} 
    onchange={loadData} />
</div>


<style>
    .task{
        display: flex;
        align-items: center;
    }
    input[type="checkbox"] {
        width: 20px;
        height: 20px;
        cursor: pointer;
    }
</style>
