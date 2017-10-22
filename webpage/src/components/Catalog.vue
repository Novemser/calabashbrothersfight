<template>
	<div class="index-page">
		<img class="logo" src="//file.baixing.net/201710/e50a70e9dc6c6a2766a84a3780db286d.png" alt="pic"/>
		<p>欢迎来到葫芦娃的王国。</p>
		<p>在这里,7只英勇的葫芦娃将与诡计多端的白骨精展开一场计算机科学的较量……</p>
		<p>亲爱的玩家，在这里，您是我们的调度者。</p>
		<p>在你面前是一系列的用<span class="high-light">Go</span>语言写的并行程序的若干个<span class="high-light">协程</span>(Goroutines)。</p>
		<p>你的目标就是使用任何方式尝试使其运行<span class="high-light">故障</span>。</p>
		<p>比如，你让两个协程同时执行到<span class="high-light">临界区</span>。(critical_section)</p>
		<p>或者，尝试到达panic("exit_here")</p>
	</div>
</template>

<script>
	export default {
		name: '',
		props: {},
		data () {
			return {}
		},
		methods: {},
		mounted: function () {
		}
	}
</script>

<style lang="stylus">
	.index-page
		font-size 16px
		.high-light
			color red
		.logo
			width 900px
</style>
