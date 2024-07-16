// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func BaseLayout(nav bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\" min-h-screen flex flex-col\"><div class=\"bg-emerald-600 flex\"><label class=\"inline-flex items-center my-2 cursor-pointer\"><input id=\"themeToggleSwicher\" type=\"checkbox\" value=\"\" class=\"sr-only peer\"><div class=\"relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:w-5 after:h-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"ms-3 text-sm font-medium text-gray-900 dark:text-gray-300\">Dark Mode</span></label><div class=\" px-2 mx-2 text-xs flex justify-between gap-2 font-medium \"><a class=\" my-auto text-orange-600 bg-orange-800\" href=\"/\">home</a> <a class=\" my-auto text-rose-600 bg-rose-800\" href=\"/login\">Login</a> <a class=\" my-auto text-violet-600 bg-violet-800\" href=\"/register\">Register</a> <a class=\" my-auto text-lime-600 bg-lime-800\" href=\"/tags\">tags</a> <a class=\" my-auto text-yellow-600 bg-yellow-800\" href=\"/home\">xd</a></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n            let themeToggleSwicher = document.getElementById(\"themeToggleSwicher\");\n            \n            const setTheme = (isDark) => {\n                if (isDark) {\n                    document.documentElement.classList.add('dark');\n                    localStorage.setItem('color-theme', 'dark');\n                    themeToggleSwicher.checked = true;\n                } else {\n                    document.documentElement.classList.remove('dark');\n                    localStorage.setItem('color-theme', 'light');\n                    themeToggleSwicher.checked = false;\n                }\n            };\n        \n            themeToggleSwicher.onchange = () => {\n                const isDark = themeToggleSwicher.checked;\n                setTheme(isDark);\n            };\n        </script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
