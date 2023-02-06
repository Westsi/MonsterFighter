<script>
  const makeNewUser = () => {
    fetch("http://localhost:8000/new/user", {
      method: 'POST',
      mode: 'no-cors',
      cache: 'no-cache',
      headers: {
          'Content-Type': 'application/json'
      },
      redirect: 'follow',
      referrerPolicy: 'no-referrer',
      body: JSON.stringify({"Name": "testUser", "Password": "testPassword", "Email": "myemail@test.com"})
    })
  }
</script>

<main>
  <button on:click={() => makeNewUser()}>Make New User</button>
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>
