<template>
	<div class="section">
		<h1 class="title">{{ label }} - {{ title }}</h1>
		<div class="panel">
			<p class="introduction" v-html="description"></p>
			<div class="source-controls">
				<mu-raised-button icon="undo" label="撤销" @click="undo()"
								  primary :disabled="!canUndo"></mu-raised-button>
				<mu-raised-button icon="refresh" label="重置" @click="reload()"
								  secondary></mu-raised-button>
			</div>
		</div>
		<p class="objective"><span>目标：</span>{{ victoryCond }}</p>
		<h2 class="subtitle" title="thread">代码</h2>
		<div class="source">
			<div v-for="(program, thread) in programs" class="thread" :key="thread">
				<h3 class="thread-header">
					协程 - {{ thread }}
				</h3>
				<div class="controls">
					<mu-raised-button icon="play_arrow" label="步进" @click="stepThread(thread)"
					  	:disabled="!program['canStepNext']"></mu-raised-button>
					<mu-raised-button icon="zoom_in" label="展开" @click="expand(thread)"
						:disabled="!program['canCurrentExpand']"></mu-raised-button>
				</div>
				<div class="code">
					<div v-for="(expression, index) in program['code']">
						<div class="instruction" :class="{ current: program['current'][0] === index }" :key="index">
							<span class="indent">{{ expression['indent'] | showTab }}</span>
							<span class="block" :title="expression['description']"
								  v-html="highlight(expression['code'], expression['name'])"></span>
						</div>
						<div class="instruction" v-if="expression['expanded']" :key="_index" style="font-size: 14px"
							 v-for="(_expression, _index) in expression['expandInstructions']"
							 :class="{ current: program['current'][1] === _index }" >
							<span class="indent">{{ expression['indent'] + 1 | showTab }}</span>
							<span class="block" :title="_expression['description']"
								  v-html="highlight(_expression['code'], _expression['name'])"></span>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div class="context">
			<h2 class="subtitle">变量</h2>
			<div class="variable" v-for="(variable, index) in context" :key="index">
				<span>{{ variable['name'] }}</span>
				<span>:=</span>
				<span>{{ variable['value'] }}</span>
			</div>
		</div>
	</div>
</template>

<script>
	import 'tippy.js/dist/tippy.css'
	import tippy from 'tippy.js/dist/tippy.min'
	import hightlightRules from '../assets/highlight-rule'
	export default {
		name: 'beginning',
		props: {},
		data () {
			return {
				level: 0,
				label: '',
				title: '',
				description: '',
				victoryCond: '',
				programs: [],
				context: [],
				gameStatus: 0,
				canUndo: false,
				_tippy: null
			}
		},
		methods: {
			fetchData (url, cb = function () {}) {
				this.axios.get(url).then((response) => {
					if (response && response.data) {
						const data = response.data
						if (data.status === 0) {
							this.load(this.setIndent(data.data), cb)
						} else {
							this.openDialog('错误', data.msg)
							console.error(data.msg)
						}
					}
				}).catch(console.error.bind(this))
			},
			setIndent (data) {
				const programs = data.programs
				for (let p = 0; p < programs.length; p++) {
					const program = programs[p]['code']
					let indent = 0
					for (let i = 0; i < program.length; i++) {
						const code = program[i]
						code.indent = indent
						if (['If statement', 'For statement'].indexOf(code.description) !== -1) {
							indent++
						} else if (['End if statement', 'End of for'].indexOf(code.description) !== -1) {
							indent--
							code.indent = indent
						}
					}
				}
				return data
			},
			mock () {
				return {
					title: '教程 1: 接口',
					description: `
						WOW你来到了这里。<br/>
						亲爱的玩家，这里，您是我们的调度者。<br/>
						在你面前是一个用<span class="high-light">Go</span>语言写的并行程序的两个
						<span class="high-light">协程</span>(Goroutines)。<br/>
						你的目标就是使用任何方式尝试使其运行<span class="high-light">故障</span>。<br/>
						而在当前关卡，你必须让两个协程同时执行到临界区。
					`,
					programs: [
						{
							current: [1, 0],
							canStepNext: true,
							canCurrentExpand: false,
							code: [
								{
									name: 'comment',
									code: '// 这里是第一个协程',
									indent: 0
								},
								{
									name: 'business',
									code: 'business_logic();',
									indent: 0
								},
								{
									name: 'critical',
									code: 'critical_section();',
									indent: 0
								},
								{
									name: 'business',
									code: 'business_logic();',
									indent: 0
								}
							]
						},
						{
							current: [1, 1],
							canStepNext: true,
							canCurrentExpand: false,
							code: [
								{
									name: 'comment',
									code: '// 这里是第二个协程',
									expanded: false,
									expandInstructions: [],
									indent: 0
								},
								{
									name: 'expression',
									code: 'a = a + 1;',
									expanded: true,
									expandInstructions: [
										{
											name: 'expression',
											code: 'temp = a;',
											expandInstructions: [],
											indent: 1
										},
										{
											name: 'expression',
											code: 'a = temp + 1;',
											expandInstructions: [],
											indent: 1
										}
									],
									indent: 0
								},
								{
									name: 'critical',
									code: 'critical_section();',
									expanded: false,
									expandInstructions: [],
									indent: 0
								},
								{
									name: 'business',
									code: 'business_logic();',
									expanded: false,
									expandInstructions: [],
									indent: 0
								}
							]
						},
						{
							current: [0, 0],
							canStepNext: true,
							canCurrentExpand: false,
							code: [
								{
									name: 'comment',
									code: '// 这里是第三个协程',
									expanded: false,
									expandInstructions: [],
									indent: 0
								},
								{
									name: 'if',
									code: 'if (flag) {',
									expanded: false,
									expandInstructions: [],
									indent: 0
								},
								{
									name: 'business',
									code: 'business_logic();',
									expanded: false,
									expandInstructions: [],
									indent: 1
								},
								{
									name: 'endif',
									code: '}',
									expanded: false,
									expandInstructions: [],
									indent: 0
								}
							]
						}
					],
					context: [
						{
							name: 'flag',
							value: true
						},
						{
							name: 'a',
							value: 0
						}
					],
					gameStatus: 0
				}
			},
			load (data, cb) {
				this.label = data.label
				this.title = data.title
				this.description = data.description
				this.victoryCond = data.victoryCond
				this.programs = data.programs
				this.context = data.context
				this.gameStatus = data.gameStatus
				this.canUndo = data.canUndo
				cb && cb()
			},
			/**
			 * 步进下一条指令
			 * @param thread 协程索引$index
			 */
			stepThread (thread) {
				const currentLine = this.programs[thread]['current'][0]
				this.fetchData(`/api/stepthread/${this.level}/${thread}/${currentLine}`)
			},
			/**
			 * 展开一条指令
			 * @param thread 协程索引$index
			 */
			expand: function (thread) {
				const currentLine = this.programs[thread]['current'][0]
				this.fetchData(`/api/expand/${this.level}/${thread}/${currentLine}`)
			},
			undo: function () {
				this.fetchData(`/api/undo/${this.level}`)
			},
			reload: function () {
				this.fetchData(`/api/level/${this.level}`, () => {
					this._tippy && this._tippy.destroyAll()
					/**
					 * hover效果
					 * 包需要在文档加载完毕之后执行
					 */
					setTimeout(() => {
						this._tippy = tippy('.block', {
							position: 'left',
							animation: 'shift',
							duration: 300,
							interactive: true,
							arrow: true
						})
					}, 100)
				})
			},
			highlight: function (code, name) {
				const rule = hightlightRules[name]
				if (!rule) return code
				if (!rule.string) return `<span style="color:${rule.color}">${code}</span>`
				if (Array.isArray(rule.string)) {
					rule.string.forEach(function (str) {
						code = code.replace(str, `<span style="color:${rule.color}">${str}</span>`)
					})
					return code
				}
				return code.replace(rule.string, `<span style="color:${rule.color}">${rule.string}</span>`)
			}
		},
		watch: {
			'$route': function () {
				this.level = this.$route.params.level
				this.reload()
			},
			'gameStatus': function () {
				switch (this.gameStatus) {
				case -1:
					this.openDialog('失败', '您太惨了2333')
					break
				case 1:
					this.openDialog('成功', 'WOW你好厉害!恭喜!')
					break
				case 0:
				default:
				}
			}
		},
		mounted: function () {
			this.level = this.$route.params.level
			this.reload()
//			this.load(this.mock())
		},
		filters: {
			showTab: function (quantity) {
				return new Array(quantity).fill('  ').join('')
			}
		}
	}
</script>

<style lang="stylus">
</style>
