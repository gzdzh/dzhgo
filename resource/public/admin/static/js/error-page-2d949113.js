import{u as s}from"./browser-5c66c964.js";import"./vue-echarts-cb4fb6c4.js";import"./element-plus-cc31d756.js";import"./store-acbd10e3.js";import"./index-95085aed.js";import{u as e}from"./index-6dccf6bb.js";import{d as r,W as o,o as t,c as a,a as i,b as u,J as n,N as c}from"./@vue.runtime-core-6173334e.js";import{r as p,u as m}from"./@vue.reactivity-54707199.js";import{L as d}from"./@vue.shared-0cc4c744.js";import{_ as l}from"./_plugin-vue_export-helper-1b428a4d.js";const f={class:"error-page"},v={class:"code"},j={class:"desc"},y={key:0,class:"btns"},_={key:1,class:"btns"},b=r({name:"undefined"}),h=l(r({...b,props:{code:Number,desc:String},setup(r){const{router:l}=e(),{user:b}=s(),h=p(!1);function k(){l.push("/login")}async function g(){h.value=!0,b.logout()}function x(){l.push("/")}return(s,e)=>{const p=o("el-button");return t(),a("div",f,[i("h1",v,d(r.code),1),i("p",j,d(r.desc),1),m(b).token||h.value?(t(),a("div",y,[u(p,{onClick:x},{default:n((()=>[c("回到首页")])),_:1}),u(p,{type:"primary",onClick:g},{default:n((()=>[c("重新登录")])),_:1})])):(t(),a("div",_,[u(p,{type:"primary",onClick:k},{default:n((()=>[c("返回登录页")])),_:1})]))])}}}),[["__scopeId","data-v-ddc7187a"]]);export{h as E};