import{d as e,q as l,W as a,o as t,c as o,E as s,b as i,J as u,N as c,n,I as p,M as d}from"./@vue.runtime-core-6173334e.js";import{r,u as m}from"./@vue.reactivity-54707199.js";import{L as f}from"./@vue.shared-0cc4c744.js";const v={class:"cl-upload-space__wrap"},_=e({name:"cl-upload-space"}),h=e({..._,props:{title:{type:String,default:"文件空间"},limit:{type:Number,default:9},accept:String,showBtn:{type:Boolean,default:!0}},emits:["confirm"],setup(e,{expose:_,emit:h}){const y=r(!1),g=r(),k=l((()=>{var e;return(null==(e=g.value)?void 0:e.selection)||[]}));function b(){y.value=!0,n((()=>{var e;null==(e=g.value)||e.clear()}))}function w(){y.value=!1}function x(){h("confirm",k.value),w()}return _({open:b,close:w}),(l,n)=>{const r=a("el-button"),_=a("cl-upload-panel"),h=a("cl-dialog");return t(),o("div",v,[s(l.$slots,"default",{},(()=>[e.showBtn?(t(),p(r,{key:0,onClick:b},{default:u((()=>[c("点击上传")])),_:1})):d("",!0)])),i(h,{modelValue:y.value,"onUpdate:modelValue":n[0]||(n[0]=e=>y.value=e),title:e.title,height:"650px",width:"1070px","keep-alive":"","custom-class":"cl-upload-space__dialog","close-on-click-modal":!1,"append-to-body":""},{footer:u((()=>[i(r,{onClick:w},{default:u((()=>[c("取消")])),_:1}),i(r,{disabled:0==m(k).length,type:"success",onClick:x},{default:u((()=>[c("选择 "+f(m(k).length)+"/"+f(e.limit),1)])),_:1},8,["disabled"])])),default:u((()=>[i(_,{limit:e.limit,accept:e.accept,ref_key:"Panel",ref:g},null,8,["limit","accept"])])),_:1},8,["modelValue","title"])])}}});export{h as default};
