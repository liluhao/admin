import{j as e,a0 as i,bi as t,K as s,o as n,m as r,P as l,n as o,q as a,X as p}from"./vendor.0af35baf.js";/* empty css              */import{_ as c}from"./PageWrapper.0f991710.js";import{_ as u,u as g}from"./useDescription.31444321.js";import{G as d,D as m,S as v}from"./siteSetting.26a5fe52.js";import"./index.36aa7878.js";import"./usePageContext.37d2295e.js";/* empty css              *//* empty css              */import"./onMountedOrActivated.28436c92.js";var f=e({name:"AboutPage",components:{Description:u,PageWrapper:c},setup(){const{pkg:e,lastBuildTime:s}={pkg:{dependencies:{"@iconify/iconify":"^2.0.1","@vueuse/core":"^5.0.2","@zxcvbn-ts/core":"^0.3.0","ant-design-vue":"2.1.2",axios:"^0.21.1","crypto-js":"^4.0.0",echarts:"^5.1.2","lodash-es":"^4.17.21",mockjs:"^1.1.0",nprogress:"^0.2.0","path-to-regexp":"^6.2.0",pinia:"2.0.0-alpha.13",qrcode:"^1.4.4",sortablejs:"^1.13.0",vue:"3.0.11","vue-i18n":"9.1.6","vue-router":"^4.0.8","vue-types":"^3.0.2"},devDependencies:{"@commitlint/cli":"^12.1.4","@commitlint/config-conventional":"^12.1.4","@iconify/json":"^1.1.354","@purge-icons/generated":"^0.7.0","@types/codemirror":"^5.60.0","@types/crypto-js":"^4.0.1","@types/fs-extra":"^9.0.11","@types/inquirer":"^7.3.1","@types/lodash-es":"^4.17.4","@types/mockjs":"^1.0.3","@types/node":"^15.12.2","@types/nprogress":"^0.2.0","@types/qrcode":"^1.4.0","@types/qs":"^6.9.6","@types/sortablejs":"^1.10.6","@typescript-eslint/eslint-plugin":"^4.26.1","@typescript-eslint/parser":"^4.26.1","@vitejs/plugin-legacy":"^1.4.1","@vitejs/plugin-vue":"^1.2.3","@vitejs/plugin-vue-jsx":"^1.1.5","@vue/compiler-sfc":"3.0.11",autoprefixer:"^10.2.6",commitizen:"^4.2.4","conventional-changelog-cli":"^2.1.1","cross-env":"^7.0.3",dotenv:"^10.0.0",eslint:"^7.28.0","eslint-config-prettier":"^8.3.0","eslint-define-config":"^1.0.8","eslint-plugin-prettier":"^3.4.0","eslint-plugin-vue":"^7.10.0",esno:"^0.7.1","fs-extra":"^10.0.0","http-server":"^0.12.3",inquirer:"^8.1.0","is-ci":"^3.0.0",less:"^4.1.1","lint-staged":"^11.0.0",postcss:"^8.3.0",prettier:"^2.3.1","pretty-quick":"^3.1.0",rimraf:"^3.0.2","rollup-plugin-visualizer":"5.5.0",stylelint:"^13.13.1","stylelint-config-prettier":"^8.0.2","stylelint-config-standard":"^22.0.0","stylelint-order":"^4.1.0","ts-node":"^10.0.0",typescript:"4.3.2",vite:"2.3.3","vite-plugin-compression":"^0.2.5","vite-plugin-html":"^2.0.7","vite-plugin-imagemin":"^0.3.2","vite-plugin-mock":"^2.7.1","vite-plugin-purge-icons":"^0.7.0","vite-plugin-pwa":"^0.7.3","vite-plugin-style-import":"^0.10.1","vite-plugin-svg-icons":"^0.7.0","vite-plugin-theme":"^0.8.1","vite-plugin-windicss":"^1.0.3","vue-eslint-parser":"^7.6.0","vue-tsc":"^0.1.7"},name:"vben-admin",version:"2.4.2"},lastBuildTime:"2021-06-14 19:18:33"},{dependencies:n,devDependencies:r,name:l,version:o}=e,a=[],p=[],c=e=>s=>i(t,{color:e},(()=>s)),u=e=>t=>i("a",{href:t,target:"_blank"},e),f=[{label:"版本",field:"version",render:c("blue")},{label:"最后编译时间",field:"lastBuildTime",render:c("blue")},{label:"文档地址",field:"doc",render:u("文档地址")},{label:"预览地址",field:"preview",render:u("预览地址")},{label:"Github",field:"github",render:u("Github")}],y={version:o,lastBuildTime:s,doc:m,preview:v,github:d};Object.keys(n).forEach((e=>{a.push({field:e,label:e})})),Object.keys(r).forEach((e=>{p.push({field:e,label:e})}));const[b]=g({title:"生产环境依赖",data:n,schema:a,column:3}),[j]=g({title:"开发环境依赖",data:r,schema:p,column:3}),[h]=g({title:"项目信息",data:y,schema:f,column:2});return{register:b,registerDev:j,infoRegister:h,name:l,GITHUB_URL:d}}});const y={class:"flex justify-between items-center"},b={class:"flex-1"},j=p(" 是一个基于Vue3.0、Vite、 Ant-Design-Vue 、TypeScript 的后台解决方案，目标是为中大型项目开发,提供现成的开箱解决方案及丰富的示例,原则上不会限制任何代码用于商用。 ");f.render=function(e,i,t,p,c,u){const g=s("Description"),d=s("PageWrapper");return n(),r(d,{title:"关于"},{headerContent:l((()=>[o("div",y,[o("span",b,[o("a",{href:e.GITHUB_URL,target:"_blank"},a(e.name),9,["href"]),j])])])),default:l((()=>[o(g,{onRegister:e.infoRegister,class:"enter-y"},null,8,["onRegister"]),o(g,{onRegister:e.register,class:"my-4 enter-y"},null,8,["onRegister"]),o(g,{onRegister:e.registerDev,class:"enter-y"},null,8,["onRegister"])])),_:1})};export default f;
