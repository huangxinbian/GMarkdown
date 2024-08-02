<template>
	<div class="page">
		<div class="mavon-pnl" :class="'module-'+ state" v-show="!isPrint">
			<mavon-editor v-model="content"
				class="md"
				ref="md"
				_props="props"
				:subfield="subfield"
				:editable="true"
				:ishljs="true"
				:toolbars-flag="true"
				:defaultOpen="defaultOpen"
				:toolbars="toolbars"
				:config="configs"
				@imgAdd="onImgAdd"
				@change="onChange"
				@previewToggle="onPreviewToggle"
				@subfieldToggle="onSubfieldToggle"
				@save="onSave"
				_externalLink="externalLink">
				<template slot="left-toolbar-before">
					<button
						type="button"
						@click="onOpen"
						class="iconfont icon-wenjianjiadakai"
						aria-hidden="true"
						title="打开文件(Alt+O)">
					</button>
					<button
						type="button"
						@click="onViewCol"
						class="iconfont icon-qiehuan"
						aria-hidden="true"
						title="模式转换(Alt+W)">
					</button>
					<button
						v-if="state != 1"
						type="button"
						@click="onEdit"
						class="iconfont icon-bianji"
						aria-hidden="true"
						title="编辑文件(Alt+E)">
					</button>
					<button
						v-if="state != 2"
						type="button"
						@click="onView"
						class="iconfont icon-liulan"
						aria-hidden="true"
						title="浏览文件(Alt+V)">
					</button>
					<div class="hr"></div>
				</template>
				<template slot="left-toolbar-after">
					<button
						type="button"
						@click="onSaveAs"
						class="iconfont icon-baocun"
						aria-hidden="true"
						title="另存文件(Alt+S)">
					</button>
					<button
						_if="state == 2"
						type="button"
						@click="onPrint"
						class="iconfont icon-dayin"
						aria-hidden="true"
						title="打印(Alt+P)">
					</button>
					<button ref="printBtn" v-show="false" v-print="'#printMe'">隐藏打印</button>
					<button
						type="button"
						@click="onAbout"
						class="iconfont icon-zhuyi"
						aria-hidden="true"
						title="关于(Alt+M)">
					</button>
				</template>
			</mavon-editor>
		</div>
		<div v-if="showAbout" class="dialog-pnl">
			<div class="shade" @click="aboutClose"></div>
			<div class="dialog">
				<div class="head not-select">
					<span class="title">关于</span>
					<span class="close iconfont icon-cuowu-fill" @click="aboutClose"></span>
				</div>
				<div class="center">
					<div class="title">
						<p class="name">GMarkdown</p>
						<p class="varsion">version: 0.0.1</p>
					</div>
					<div class="row">
						<span class="not-select">软件作者：</span>
						<span>黄新变</span>
					</div>
					<div class="row">
						<span class="not-select">联系电话：</span>
						<span>15077120431</span>
					</div>
					<div class="row">
						<span class="not-select">联系邮箱：</span>
						<span>359664529@qq.com</span>
					</div>
					<div class="row">
						<span class="not-select">友情赞助：</span>
						<img :src="weichatImg">
					</div>
				</div>
				<div class="bottom not-select">
					<button ref="aboutBtn" class="btn" @click="aboutClose">关闭</button>
				</div>
			</div>
		</div>
		<div v-if="showMsg" class="dialog-pnl">
			<div class="shade" @click="msgClose"></div>
			<div class="dialog">
				<div class="head not-select">
					<span class="title">{{msg.sTitle}}</span>
					<span class="close iconfont icon-cuowu-fill" @click="msgClose"></span>
				</div>
				<div class="center">
					<span class="msg">{{msg.sMsg}}</span>
				</div>
				<div class="bottom not-select">
					<button ref="msgBtn" class="btn" @click="msgClose">确定</button>
				</div>
			</div>
		</div>
		<div id="printMe" v-show="isPrint">
			<mavon-editor
				ref="mdPrint"
				v-model="printHtml"
				:editable="false"
				:ishljs="true"
				defaultOpen="preview"
				:shortCut="false"
				:subfield="false"
				@change="mdPrintChange"
				box-shadow-style="0 0 0 0 rgba(0,0,0,0)"
				:toolbars-flag="false"/>
		</div>
	</div>
</template>

<script>
	import "mavon-editor/dist/css/index.css"

	import "highlight.js/styles/googlecode.css"
	//import hljs from "highlight.js" //导入代码高亮文件

	export default {
		data() {
			return {
				isPrint: false,
				showAbout: false,
				showMsg: false,
				msg: {
					sTitle: "提示",
					sMsg: ""
				},
				nowFile: "",
				content: "",
				contentOld: "",
				printHtml: "",
				configs: {},
				state: 3,//1:edit;2:view;3:endAndView
				defaultOpen: "edit",
				subfield: true,
				externalLink: {
					markdown_css: function() {
						// 这是你的markdown css文件路径
						return "/mavon-editor/markdown/github-markdown.min.css";
					},
					hljs_js: function() {
						// 这是你的hljs文件路径
						return "/mavon-editor/highlightjs/highlight.min.js";
					},
					hljs_css: function(css) {
						// 这是你的代码高亮配色文件路径
						return "/mavon-editor/highlightjs/styles/" + css + ".min.css";
					},
					hljs_lang: function(lang) {
						// 这是你的代码高亮语言解析路径
						return "/mavon-editor/highlightjs/languages/" + lang + ".min.js";
					},
					katex_css: function() {
						// 这是你的katex配色方案路径路径
						return "/mavon-editor/katex/katex.min.css";
					},
					katex_js: function() {
						// 这是你的katex.js路径
						return "/mavon-editor/katex/katex.min.js";
					},
				},
				toolbars: {},
				viewToolbars: {
					readmodel: true, // 沉浸式阅读
					help: true, // 帮助
					navigation: true, // 导航目录
				},
				editToolbars: {
					bold: true, // 粗体
					italic: true, // 斜体
					header: true, // 标题
					underline: true, // 下划线
					strikethrough: true, // 中划线
					mark: true, // 标记
					superscript: true, // 上角标
					subscript: true, // 下角标
					quote: true, // 引用
					ol: true, // 有序列表
					ul: true, // 无序列表
					link: true, // 链接
					imagelink: true, // 图片链接
					code: true, // code
					table: true, // 表格
					fullscreen: false, // 全屏编辑
					/* 1.3.5 */
					undo: true, // 上一步
					redo: true, // 下一步
					trash: true, // 清空
					save: true, // 保存（触发events中的save事件）
					/* 1.4.2 */
					navigation: false, // 导航目录
					/* 2.1.8 */
					alignleft: true, // 左对齐
					aligncenter: true, // 居中
					alignright: true, // 右对齐
					readmodel: false, // 沉浸式阅读
					htmlcode: false, // 展示html源码
					help: false, // 帮助
					subfield: false, // 单双栏模式
					preview: false, // 预览
				},
				resizeObserver: null,
				observer: null,
				lineDivs: [],
				weichatImg: "data:image/png;base64,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"
			}
		},
		methods: {
			init() {
				//this.changeState(3);
				this.loadInitFile();
			},
			
			loadInitFile() {
				if (!this.go.ReadNowFile) {
					return;
				}
				
				this.go.ReadNowFile().then((res)=>{
					if (!res.success) {
						this.changeState(3);
						return;
					}
					
					this.changeState(2);
					this.nowFile = res.msg;
					this.content = res.data;
					
					if ((this.content != "") && (this.content.includes("```"))) {
						this.$nextTick(() => {
							const timer = setTimeout(() => {
								this.addLineNumber(this.$refs.md.$el);
								this.addEditLineNumber(this.$refs.md.$el);
								clearTimeout(timer);
							}, 3000);
						});
					}
				});
			},
			
			onOpen() {
				this.go.OpenFile(this.getFilePath(this.nowFile)).then((res) => {
					if (!res.success) {
						if (res.msg == "文件路径为空") {
							return;
						}
						this.onMsg("错误", res.msg);
						//this.content = res.msg;
						return;
					}
					
					this.nowFile = res.msg;
					this.$nextTick(() => {
						this.content = res.data;
					});
					
					this.$nextTick(() => {
						this.changeState(2);
					});
				});
			},
			
			onEdit() {
				this.changeState(1);
			},
			
			onView() {
				this.changeState(2);
			},
			
			onViewCol() {
				let temp = this.state +1;//1:edit;2:view;3:endAndView
				if (temp > 3) {
					temp = 1;
				}
				
				this.changeState(temp);
			},
			
			changeState(iState) {
				this.state = iState;
				
				switch (this.state){
					case 1:
						this.defaultOpen = "edit";
						this.toolbars = this.editToolbars;
						this.subfield = false;
						break;
					case 2:
						this.defaultOpen = "preview";
						this.toolbars = this.viewToolbars;
						this.subfield = false;
						break;
					case 3:
						this.defaultOpen = "preview";
						this.toolbars = this.editToolbars;
						this.subfield = true;
						break;
					default:
						break;
				}
				
				this.$nextTick(() => {
					this.addLineNumber(this.$refs.md.$el);
					this.addEditLineNumber(this.$refs.md.$el);
				});
			},

			/**
			 * 保存
			 */
			onSave() {
				this.go.SaveFile(this.nowFile, this.content).then((res) => {
					if (!res.success) {
						this.onMsg("错误", res.msg);
						// this.go.ErrorMsg("错误", res.msg);
						return;
					}
					
					//this.content = res.data;//返回的是路径
				});
			},
			
			// 内容改变事件
			onChange(value, render) {
				this.$nextTick(() => {
					this.addEditLineNumber(this.$refs.md.$el);
					this.addLineNumber(this.$refs.md.$el);
				});
				// this.sTitle = this.sTitle.trim();
				// if (this.sTitle != "") {
				// 	return;
				// }
				
				// let temp = value.trim();
				// if (temp == "") {
				// 	return;
				// }
				
				// let iEd = temp.indexOf("\n");
				// temp = temp.substring(0, iEd).trim();
				
				// while(temp.indexOf("#") == 0) {
				// 	temp = temp.substring(1).trim();
				// }
				
				// this.sTitle = temp;
			},
			
			onPreviewToggle(status, value) {
				this.$nextTick(() => {
					this.addLineNumber(this.$refs.md.$el);
				});
			},
			
			onSubfieldToggle(status, value) {
				this.$nextTick(() => {
					this.addLineNumber(this.$refs.md.$el);
				});
			},
			
			/**
			 * 另存
			 */
			onSaveAs() {
				this.go.SaveAsFile(this.getFilePath(this.nowFile), this.content).then((res) => {
					if (!res.success) {
						if (res.data == 8001) {
							return;
						}
						
						if (res.msg.includes("The system cannot find the file specified")) {
							return;
						}
						
						this.onMsg("错误", res.msg);
						//this.go.ErrorMsg("错误", res.msg);
						//this.content = res.msg;
						return;
					}
					
					this.nowFile = res.data;//返回的是路径
				});
			},
			
			onImgAdd(pos, $file) {
				this.$refs.md.$img2Url(pos, $file.miniurl);//不使用图床的做法
				return;
			},
			
			getFilePath(path) {
				let result = path;
				if (result == "") {
					return result;
				}
				
				result = result.replace(/\\/g,"/");
				let i = result.lastIndexOf("/");
				if (i < 0) {
					return "";
				}
				
				result = result.substring(0, i);
				
				return result;
			},
			
			onAbout() {
				this.showAbout = true;
				this.$nextTick(() => {
					this.$refs.aboutBtn.focus();
				});
			},
			
			aboutClose() {
				this.showAbout = false;
				this.$nextTick(() => {
					this.$refs.md.focus();
				});
			},
			
			onPrint() {
				//this.go.WindowPrint();//只能打印窗口
				//this.printHtml = this.content + "\r\n<br /><br /><br /><br /><br /><br />";//补偿缺失的行
				if (this.printHtml != this.content) {
					this.printHtml = this.content;
					this.$nextTick(()=>{//这行不能少
						this.addLineNumberToPrint(this.$refs.mdPrint.$el);//这个方式行号会有问题,行号必须使用 position: absolute;来进行布局,否则会乱
					});
				}
				
				this.isPrint = true;
				
				this.$nextTick(()=>{
					const btn = this.$refs.printBtn;//this.$refs.printBtn.$el
					btn.click();//异步方式
					this.isPrint = false;
					
					//这个方式行号没问题,但分页有问题
					//this.addLineNumberToPrint(this.$refs.mdPrint.$el);
					// this.$nextTick(()=>{
					// 	window.print();//会进入等待
					// 	this.isPrint = false;
					// })
				});
			},
			
			mdPrintChange() {
				// this.$nextTick(() => {
				// 	this.addLineNumber(this.$refs.mdPrint.$el);
				// });
				// if (this.printHtml == "") {
				// 	return;
				// }
			},
			
			onMsg(sTitle, sMsg) {
				this.msg = {
					sTitle: sTitle,
					sMsg: sMsg
				};
				
				this.showMsg = true;
				this.$nextTick(() => {
					this.$refs.msgBtn.focus();
				});
			},
			
			msgClose() {
				this.showMsg = false;
				this.$nextTick(() => {
					this.$refs.md.focus();
				});
			},
			
			/**
			 * 添加代码行号
			 */
			addLineNumber(mdElement) {
				if ((this.content == "") || (!this.content.includes("```"))) {
					return;
				}
				
				const hljsList = mdElement.querySelectorAll(".v-note-show pre .hljs");
				if (!hljsList) {
					return;
				}
				
				for (let k = 0; k < hljsList.length; k++) {
					const codeDiv = hljsList[k].firstChild;
					
					let lineDiv = hljsList[k].querySelector(".line-div")
					if (!lineDiv) {
						lineDiv = document.createElement("div");
						lineDiv.className = "line-div";
						hljsList[k].insertBefore(lineDiv, codeDiv);
						
						this.observeElementRemoval(lineDiv);//加入监听
						let resizeObserver = new ResizeObserver((entries) => {
							this.addLineNumber2(lineDiv);
						});
						resizeObserver.observe(lineDiv);
						
						let has = false;
						for (let i = 0; i < this.lineDivs.length; i++) {
							if (this.lineDivs[i].id == lineDiv) {
								has = true;
								break;
							}
						}
						
						if (!has) {
							this.lineDivs.push({id: lineDiv, resizeObj: resizeObserver});
						}
					}
				}
			},
			
			addLineNumber2(lineDiv) {
				this.$nextTick(()=>{
					let iSum = Math.round(lineDiv.offsetHeight / parseFloat(getComputedStyle(lineDiv).lineHeight));
					
					let temp = "";
					for (let i = 0; i < iSum ; i++) {
						temp = temp + "<p>" + (i + 1) + "</p>";
					}
					
					lineDiv.innerHTML = temp;
				})
			},
			
			/**
			 * 添加代码行号
			 */
			addLineNumberToPrint(mdElement) {
				const hljsList = mdElement.querySelectorAll(".v-note-show pre .hljs");
				if (!hljsList) {
					return;
				}
				
				for (let k = 0; k < hljsList.length; k++) {
					const codeDiv = hljsList[k].firstChild;
					
					let lineDiv = hljsList[k].querySelector(".line-div")
					if (!lineDiv) {
						lineDiv = document.createElement("div");
						lineDiv.className = "line-div";
						hljsList[k].insertBefore(lineDiv, codeDiv);
						
						this.addLineNumber2(lineDiv);
					}
				}
			},

			/**
			 * 添加编辑区行号
			 */
			addEditLineNumber(mdElement) {
				if (this.content == "") {
					return;
				}
				
				const contentInput = mdElement.querySelector(".v-note-edit .content-input-wrapper .content-input");
				if (!contentInput) {
					return;
				}
				
				const textarea = contentInput.querySelector("textarea");
				
				let lineDiv = contentInput.querySelector(".edit-line-div")
				if (!lineDiv) {
					lineDiv = document.createElement("div");
					lineDiv.className = "edit-line-div";
					//contentInput.insertBefore(lineDiv, codeDiv);
					contentInput.appendChild(lineDiv);
				}

				lineDiv.style.height = textarea.offsetHeight + "px";
				
				this.$nextTick(()=>{
					//const iSum = Math.round(lineDiv.offsetHeight / parseFloat(getComputedStyle(lineDiv).lineHeight));
					const iSum = textarea.value.split(/\r\n|\n/).length;

					let temp = "";
					for (let i = 0; i < iSum ; i++) {
						temp = temp + "<p>" + (i + 1) + "</p>";
					}
					
					lineDiv.innerHTML = temp;
					//--设置宽度--//
					let iWidth = (iSum + "").length * 14 + 4;
					lineDiv.style.width = iWidth + "px";

					let iMarginLeft = iWidth - 25;//25是本身框体左边距
					textarea.style.marginLeft = (iMarginLeft > 0 ? iMarginLeft : 0) + "px";//如果行宽未超出偏移值,则保持不变
				});
			},
			
			observeElementRemoval(lineDiv) {
				this.observer = new MutationObserver((mutationsList) => {
					mutationsList.forEach(mutation => {
						if ((mutation.type == 'childList') && (mutation.target == lineDiv) && (mutation.removedNodes.length > 0)) {
							//console.log('元素已被移除');

							for (let i = this.lineDivs.length -1; i >= 0; i--) {
								if (mutation.target == this.lineDivs[i].id) {
									this.lineDivs[i].resizeObj.unobserve(mutation.target);//结束对指定 Element 的监听。
									this.lineDivs[i].resizeObj.disconnect(this.lineDivs[0]);//取消特定观察者目标上所有对 Element 的监听
									this.lineDivs[i].id = null;
									this.lineDivs[i].resizeObj = null;
									this.lineDivs.splice(i, 1);
								}
								
								//this.resizeObserver.unobserve(this.lineDivs[0]);//结束对指定 Element 的监听。
								//this.resizeObserver.disconnect(this.lineDivs[0]);//取消特定观察者目标上所有对 Element 的监听
							}
						}
					});
				});
			  
				// 开始观察lineDiv的子节点变化
				this.observer.observe(lineDiv, { childList: true });
			},
			
			// onTest() {
			// 	this.msg = {
			// 		sTitle: "错误",
			// 		sMsg: "aaaaaaaaaaaaaaaaa"
			// 	};
				
			// 	this.showMsg = true;
			// },
			
			/**
			 * 快捷键事件
			 * @param {Object} event
			 */
			keyFunc(event) {
				let shortcutEvent = [				//快捷键对应的事件
					{key: "P", alt: true, fun: this.onPrint},
					{key: "E", alt: true, fun: this.onEdit},
					{key: "V", alt: true, fun: this.onView},
					{key: "W", alt: true, fun: this.onViewCol},
					{key: "S", alt: true, fun: this.onSaveAs},
					{key: "O", alt: true, fun: this.onOpen},
					{key: "M", alt: true, fun: this.onAbout}
				];
				
				this.keyDown(event, shortcutEvent);
			}
		},
		
		/**
		 * 加载后执行
		 */
		mounted() {
			this.init();
			window.addEventListener("keydown", this.keyFunc);
			window.addEventListener("keyup", this.keyup);
			
			this.resizeObserver = new ResizeObserver((entries) => {
				for (let i = 0; i < this.lineDivs.length; i++) {
					this.addLineNumber2(this.lineDivs[0]);//this.addLineNumber2(lineDiv)
				}
			});
		},
		
		//--释放前执行--//
		beforeDestroy() {
			//手动销毁监听
			window.removeEventListener("keydown", this.keyFunc);
			window.removeEventListener("keyup", this.keyup);
			//observer.disconnect();
			
			if (this.observer) {
				this.observer.disconnect();
			}
			
			for (let i = this.lineDivs.length -1; i >= 0; i--) {
				this.lineDivs[i].resizeObj.unobserve(this.lineDivs[i].id);//结束对指定 Element 的监听。
				this.resizeObserver.disconnect(this.lineDivs[i].id);//取消特定观察者目标上所有对 Element 的监听
			}
			
			this.lineDivs = [];
		}
	}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	.page {
		width: 100%;
		height: 100%;
		/* overflow: hidden; */
		background-color: #FFFFFF;
	}

	.page>.mavon-pnl {
		width: 100%;
		height: 100%;
		overflow: hidden;
	}

	.page>.mavon-pnl>.markdown-body.md {
		width: 100%;
		height: 100%;
	}
	
	.v-note-op button.iconfont {
		border: 0px;
		background-color: #ffffff;
		font-size: 13px;
		margin: 6px 0 5px 0px;
		font-size: 14px;
		padding: 4.5px 6px 5px 3.5px;
		line-height: 1;
		border-radius: 5px;
		box-sizing: border-box;
		display: inline-block;
		cursor: pointer;
		height: 28px;
		width: 28px;
	}
	
	.v-note-op button.iconfont:hover {
		background-color: #eaeaea;
	}
	
	.v-note-op div.hr{
		border-left: 1px solid #e5e5e5;
		display: inline;
		line-height: 28px;
	}
	
	.dialog-pnl {
		position: fixed;
		top: 0;
		left: 0;
		height: 100%;
		width: 100%;
		overflow: hidden;
		z-index: 9999;
	}
	
	.dialog-pnl>.shade {
		height: 100%;
		width: 100%;
		z-index: 9998;
		background-color: #00000080;
	}
	
	.dialog-pnl>.dialog {
		position: absolute;
		top: 40%;
		left: 50%;
		width: 400px; 
		transform: translate(-50%, -50%);
		background-color: #fff;
		z-index: 9999;
		border-radius: 10px;
	}
	
	.dialog-pnl>.dialog>.head {
		width: 100%;
		padding: 10px;
		box-sizing: border-box;
		display: inline-flex;
		justify-content: space-between;
		align-items: center;
	}
	
	.dialog-pnl>.dialog>.head>.title {
		padding: 0 10px 0 20px;
		font-weight: 600;
	}
	
	.dialog-pnl>.dialog>.head>.close {
		padding-right: 10px;
		cursor: pointer;
	}
	
	.dialog-pnl>.dialog>.head>.close:hover {
		font-weight: 600;
		color: #F44336;
	}
	
	.dialog-pnl>.dialog>.center {
		width: 100%;
		min-width: 250px;
		min-height: 100px;
		padding: 10px 10px;
		box-sizing: border-box;
		display: inline-flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		background-color: #e2e0e0;
	}
	
	.dialog-pnl>.dialog>.center>.title {
		width: 100%;
		text-align: center;
	}
	
	.dialog-pnl>.dialog>.center>.title>.name {
		width: 100%;
		font-size: 18px;
		font-weight: 700;
		margin-top: 4px;
		margin-bottom: 0px;
	}
	
	.dialog-pnl>.dialog>.center>.title>.varsion {
		width: 100%;
		font-size: 14px;
		margin-top: 1px;
	}
	
	.dialog-pnl>.dialog>.center>.row {
		width: 70%;
		display: inline-flex;
		justify-content: flex-start;
		padding: 4px 0px 4px 10px;
	}
	
	.dialog-pnl>.dialog>.center>.msg {
		width: 70%;
		height: 60px;/*控制div的高度*/
		box-sizing: border-box;
		white-space: normal;
		display: -webkit-box;
		-webkit-line-clamp: 3;/*控制文字行数*/
		word-break: break-all;
		text-overflow: ellipsis;
		overflow-y: hidden;
	}
	
	.dialog-pnl>.dialog>.bottom {
		width: 100%;
		height: 60px;
		padding: 5px;
		box-sizing: border-box;
		display: inline-flex;
		justify-content: center;
		align-items: center;
	}
	
	.dialog-pnl>.dialog>.bottom>.btn {
		margin: 0 5px;
		height: 36px;
		width: 76px;
		cursor: pointer;
	}

	.page>.mavon-pnl.module-1>>>.v-right-item.transition,
	.page>.mavon-pnl.module-3>>>.v-right-item.transition {
		display: none;
	}
	
	.page>.mavon-pnl>>>.v-note-edit .content-input-wrapper  {
		/* padding: 0px 25px 15px 0px !important; */
	}
	
	.page>.mavon-pnl>>>.v-note-edit .content-input-wrapper .content-input {
		position: relative;
	}
	
	.page>.mavon-pnl>>>.v-note-edit textarea {
		white-space: nowrap;
		padding-left: 10px;
		margin-left: 20px;
	}
	
	.page>.mavon-pnl>>>.v-note-edit .content-input-wrapper .edit-line-div {
		position: absolute;
		top: 0;
		left: -25px;
		width: 40px;
		height: 100%;
		padding-top: 0em;
		padding-bottom: .2em;
		margin: 0 10px 0 0;
		
		/* background-color: #dcdfe6; */
		box-sizing: border-box;
		overflow-y: clip;
		border-right: 1px solid #607D8B;
		text-align: right;
		
		-webkit-user-select: none;/*webkit浏览器*/ 
		-moz-user-select: none;/*火狐*/ 
		-ms-user-select: none;/*IE10*/ 
		user-select: none;
		color: #F44336;
	}
	
	.page>.mavon-pnl>>>.v-note-edit .content-input-wrapper .edit-line-div>p {
		width: 100%;
		text-align: right;
		margin: 0;
		padding-right: 8px;
		font-size: inherit;
		line-height: inherit;
	}
	
	.page>.mavon-pnl>>>.v-note-show pre>.hljs {
		display: flex;
	}
	
	.page>.mavon-pnl>>>.v-note-show pre>.hljs .line-div {
		width: 40px;
		height: auto;
		padding-top: .2em;
		padding-bottom: .2em;
		margin: 0 10px 0 0;
		font-size: 85%;
		background-color: #dcdfe6;
		box-sizing: border-box;
		overflow-y: clip;
		border-right: 1px solid #607D8B;
		
		-webkit-user-select: none;/*webkit浏览器*/ 
		-moz-user-select: none;/*火狐*/ 
		-ms-user-select: none;/*IE10*/ 
		user-select: none;
	}
	
	.page>.mavon-pnl>>>.v-note-show pre {
		padding: 0px !important;
	}
	
	.page>.mavon-pnl>>>.v-note-show pre>.hljs .line-div>p {
		width: 40px;
		text-align: right;
		margin: 0;
		padding-right: 8px;
		line-height: inherit;
	}
	
	#printMe>>>.v-note-show pre {
		padding: 0px !important;
	}
	
	#printMe>>>.v-note-show pre>.hljs {
		position: relative;
		padding-left: 50px;
	}
	
	#printMe>>>.v-note-show pre>.hljs .line-div {
		position: absolute;
		top: 0;
		left: 0;
		width: 40px;
		height: 100%;
		padding-top: .2em;
		padding-bottom: .2em;
		margin: 0;
		font-size: 85%;
		background-color: #dcdfe6;
		box-sizing: border-box;
		border-right: 1px solid #607D8B;
		overflow-y: clip;
		
		-webkit-user-select: none;/*webkit浏览器*/ 
		-moz-user-select: none;/*火狐*/ 
		-ms-user-select: none;/*IE10*/ 
		user-select: none;
		
		line-height: 1.7;
	}
	
	#printMe>>>.v-note-show pre>.hljs .line-div>p {
		width: 40px;
		text-align: right;
		margin: 0;
		padding-right: 8px;
		line-height: inherit;
	}
	
	@page {
		size: auto;
		/* margin: 0mm; */ /* 页边距 */
	}

	.markdown-body {
		tab-size: 4; /* \t长度 或者使用 '4em' 如果需要更精确的控制 */
	}

</style>