<template>
	<div class="hello">
		<MonacoEditor
			height="600"
			language="go"
			:code="code"
			:editorOptions="options"
			@mounted="onEditorMounted"
			@codeChange="onCodeChange"
		>
		</MonacoEditor>
	</div>
</template>

<script>
	import MonacoEditor from 'vue-monaco-editor'
	export default {
		name: 'HelloWorld',
		components: {
			MonacoEditor
		},
		data () {
			return {
				code: '// Typed away! \n',
				options: {
					selectOnLineNumbers: false
				},
				ws: null
			}
		},
		methods: {
			onEditorMounted (editor) {
				this.editor = editor
			},
			onCodeChange (editor) {
//				console.log(editor.getValue())
				this.ws.send(editor.getValue())
			}
		},
		mounted: function () {
			// this.openDialog('haha', '233')

			this.ws = new WebSocket('ws://localhost:9544/ws')
			this.ws.onopen = (event) => {
				this.ws.send('hello world~')
			}
			this.ws.onmessage = (event) => {
				console.log(event.data)
			}
		}
	}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="stylus">
	.hello
		background #fff
	.monaco-editor
		text-align left
</style>
