<script lang="ts">
import { marked } from "marked";
import { onMount } from "svelte";

export let id: string,
title = "",
from: string,
markdown = ''

onMount(async () => {

	const r = await fetch("/api/blog?q=blog&v=" + id.toString());
	try {
		const j = await r.json();
		title = j["title"][0];
		markdown = j["md"][0]

		document.getElementById("blog-content").innerHTML = marked.parse(markdown);
	} catch {
		alert("An error has occured");
	}
});
</script>

<div class="cc">
	<div class="tit">
		{title}
	</div>
	<div id="blog-content" />
</div>

<style lang="scss">
.cc {
	padding-top: 5rem;
	width: 100%;
	height: 100%;
}

.tit {
	font-size: 4rem;
	font-weight: 500;
}
</style>
