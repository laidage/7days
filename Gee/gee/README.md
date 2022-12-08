day2: 使用context（上下文）保存信息，提高用户友好度

engine——>router——>root(node)

router是用来管理路由的，这样对功能进行了分离。路由的底层原理是trim，即用node组成的树，需要实现添加和查找功能。

engine和context的联系比较小，context用来处理一些额外工作，比如设置header这些。这样用户不用进行重复操作，每个handler写一次

nodes使用trim来组织，顶部是root，root也是一个node

一个是添加路由，另一个是通过浏览器访问路由，需要分开才能理解

group分组，engine下有多个group，且拥有group的所有功能，所有group指向唯一的engine