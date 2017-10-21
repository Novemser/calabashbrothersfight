<template>
	<mu-dialog :open="status" :title="title" @close="close">
		{{ message }}
		<mu-flat-button label="确定" slot="actions" primary @click="close"/>
	</mu-dialog>
</template>

<script>
	export default {
		name: 'AlertDialog',
		data () {
			return {
				title: '',
				message: '',
				status: false
			}
		},
		methods: {
			open (title, message) {
				this.title = title
				this.message = message
				this.status = true
			},
			close () {
				this.title = ''
				this.message = ''
				this.status = false
			}
		}
	}
</script>
