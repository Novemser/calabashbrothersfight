import axios from 'axios'
import qs from 'qs'

var _dialog = null

export default {
	mountDialog (el) {
		_dialog = el
	},
	openDialog (title, message) {
		if (_dialog) {
			_dialog.open(title, message)
		}
	},
	install (Vue, options) {
		Vue.prototype.axios = axios
		Vue.prototype.qs = qs
		Vue.prototype.mountDialog = this.mountDialog
		Vue.prototype.openDialog = this.openDialog
	}
}
