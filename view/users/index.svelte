<script>
  export let users = [];
  import Header from "../Header.svelte";
</script>

<Header />
<h1>Hello users</h1>

<table border="1" cellpadding="10">
  {#if users.length > 0}
    <thead>
      {#each Object.keys(users[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each users as user}
    <tr>
      {#each Object.keys(user) as key}
        {#if key === "id"}
          <td><a href={`/${user[key]}`}>{user[key]}</a></td>
        {:else}
          <td>{user[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<style>
  table {
    border-collapse: collapse;
  }
</style>
