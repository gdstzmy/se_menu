深入理解callback函数
---

### 基本数据结构
**`DataNode`存储节点信息：**
```C++
typedef struct DataNode
{
    tLinkTableNode head;
    char*   cmd;
    char*   desc;
    int     (*handler)();
} tDataNode;
```
**`LinkTableNode`管理指针：**
```C++
struct LinkTableNode
{
    struct LinkTableNode * pNext;
};
```
**`LinkTable`是链表头指针：**
```C++
struct LinkTable
{
    struct LinkTableNode *pHead;
    struct LinkTableNode *pTail;
    int			SumOfNode;
    pthread_mutex_t mutex;

};
```
---

### lab5.2 代码工作机制：
main函数等待用户输入命令，当用户输入命令后，程序会调用`FindCmd`函数来查找对应的命令数据节点。如果找到了数据节点，则输出该命令的描述并且执行对应的操作。
```c++
printf("%s - %s\n", p->cmd, p->desc);
        if(p->handler != NULL)
        { 
            p->handler();
        }
```
----
### callback函数探幽：
其中`FindCmd`调用`SearchLinkTableNode`,`SearchLinkTableNode`函数中存在一个作为参数的函数`Condition`。
```c++
tLinkTableNode * SearchLinkTableNode(tLinkTable *pLinkTable, 
int Condition(tLinkTableNode * pNode, void * args),void * args)
```
在`menu`中存在其对应的函数实现`SearchCondition`。`*pNode`是链表管理的指针，`SearchCondition`将其转换为带有信息的`DataNode`节点。
```c++
tDataNode * pNode = (tDataNode*)GetLinkTableHead(head);
```
再通过`strcmp`进行比对以决定返回`SUCCESS`或者是`FAILURE`。
```c++
int SearchCondition(tLinkTableNode * pLinkTableNode, void * args)
{
    char * cmd = (char*) args;
    tDataNode * pNode = (tDataNode *)pLinkTableNode;
    if(strcmp(pNode->cmd, cmd) == 0)
    {
        return  SUCCESS;  
    }
    return FAILURE;	       
}
```
最后`SearchLinkTableNode`会根据`SUCCESS`或者`FAILURE`返回正确的`pNode`或者`NULL`给`FindCmd`最后返回给`main`。
```c++
    while(pNode != NULL)
    {    
        if(Condition(pNode, args) == SUCCESS)
        {
            return pNode;				    
        }
        pNode = pNode->pNext;
    }
    return NULL;
```
----
### 抽象分层
代码主要实现了`DataNode`和`LinkTableNode`管理的分流，使得节点和指针的管理分开，不会互相干扰和影响。

