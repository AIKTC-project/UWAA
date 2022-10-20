<script lang="ts">
	import { setRedOutline } from "../ts/setRedOutline";
	import Button from "../components/Button.svelte";

	let title = '', markdown = ''

	const submit = async () => {
		if (title.trim().length < 1) {
			setRedOutline(document.getElementById("title"));
			return;
		}
		if (markdown.trim().length < 1) {
			setRedOutline(document.getElementById("markdown"));
			return;
		}

		const r = await fetch('/api/blog', {method: 'POST', body: JSON.stringify({md: markdown.trim(), title: title.trim()})})
			try {
				const j = await r.json()
				if (j['_status'] === 'sfl') {
					alert('Successfully posted')
					title	= ''
					markdown = ''
				}
			} catch {
				alert('An error has occured')
			}
	};

	/*if (update) {
		const r = await fetch("/api/blog?q=blog&v=" + id.toString());
	try {
		const j = await r.json();
		title = j["title"][0];
		markdown = j["md"][0];
	} catch {
		alert("An error has occured");
	}
	}*/
</script>

<div class="con">
	<input
		id="title"
		class="tt"
		type="text"
		placeholder="Enter title for blog"
		bind:value={title}
	/>
	<textarea
		id="markdown"
		class="content"
		placeholder="Write markdown blog"
		bind:value={markdown}
	/>
	<Button callback={submit}>Post</Button>
</div>

<style lang="scss">
	.con {
		padding-top: 5rem;
		width: 90%;
	}

	.tt {
		padding: 0.5rem;
		font-size: 2rem;
	}

	.content {
		margin-top: 0.5rem;
		font-family: monospace;
		width: 100%;
		height: 30rem;
	}

	:global(code) {
		display: flex;
		padding: 0.5rem;
		margin: 1rem;
		border: 1px solid grey;
		font-family: monospace;
		color: #fff;
		background-color: #1e1e1e;
	}
</style>
