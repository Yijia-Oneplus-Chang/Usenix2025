# txt_dir = './txt//'
# if not os.path.exists(txt_dir):
#     os.makedirs(txt_dir)


# file_name = "acc_.txt"
# e_txt = open(txt_dir+file_name, 'a')
# e_txt.write("attacker 5 time:{0}s    {1}    {2}".format(execution_time,args.model,args.dataset)+"\n")
# e_txt.close()
import numpy as np
import matplotlib as mpl
import matplotlib.pyplot as plt
from matplotlib.patches import Patch
from matplotlib.lines import Line2D
import os



x_max = 930
x_em = 100
a1 = "Attack 2"
a2 = "Attack 3"
a5 = "Attack 4"
a6 = "Attack 1"
a0 = "Honest Training"
y_max =0.4
# figsize_x = 7
# figsize_y =5
# f_s = 14
# figsize_x = 7.5
# figsize_y =5
# f_s = 16
# f_w = 'normal'
# figsize_x = 7.8
# figsize_y =5
# f_s = 17
# f_w = 'normal'
# figsize_x = 7.8
# figsize_y =5
# f_s = 17
# f_w = 'normal'
# figsize_x = 8
# figsize_y =5
# f_s = 16
# f_w = 'heavy'
figsize_x = 8.5
figsize_y =5
f_s = 17
f_w = 'heavy'

current_path = os.getcwd()
# print(current_path)

lines_ele_list = []
# with open("/Users/good/Documents/IPP/code/acc__add syn100.txt", "r") as file:
with open("/Users/good/Documents/IPP/code/acc_convnetd2_cifar10.txt", "r") as file:
# with open("/Users/good/Documents/IPP/code/acc_convnetd4_cifar10.txt", "r") as file:
    content = file.read()
lines = content.split("\n")
for i in lines:
    temp = i.split("    ")
    lines_ele_list.append(temp)

lines_ele_list.pop()
# print(lines_ele_list)

cifar10 = []
cifar100 = []

# ['0.11405000000000001', 'ResNet34', 'CIFAR10', '10', 'attacker 1']
#     0                        1           2       3        4
for i in lines_ele_list:
    if i[2]=='CIFAR10':
        cifar10.append(i)
    elif i[2]=='CIFAR100':
        cifar100.append(i)

cifar10_resnet18 = []
cifar10_resnet34 = []
cifar10_convnet = []
cifar10_convnetd2 = []
cifar10_convnetd4 = []

cifar100_resnet18 = []
cifar100_resnet34 = []
cifar100_convnet = []
cifar100_convnetd2 = []
cifar100_convnetd4 = []

for i in cifar10:
    if i[1]=='ResNet18':
        cifar10_resnet18.append(i)
    elif i[1]=='ResNet34':
        cifar10_resnet34.append(i)
    elif i[1]=='ConvNet':
        cifar10_convnet.append(i)
    elif i[1]=='ConvNetD2':
        cifar10_convnetd2.append(i)
    elif i[1]=='ConvNetD4':
        cifar10_convnetd4.append(i)   

for i in cifar100:
    if i[1]=='ResNet18':
        cifar100_resnet18.append(i)
    elif i[1]=='ResNet34':
        cifar100_resnet34.append(i)
    elif i[1]=='ConvNet':
        cifar100_convnet.append(i)
    elif i[1]=='ConvNetD2':
        cifar100_convnetd2.append(i) 
    elif i[1]=='ConvNetD4':
        cifar100_convnetd4.append(i)






















# region cifar10 resnet18
cifar10_resnet18_1 = []
cifar10_resnet18_2 = []
cifar10_resnet18_5 = []
cifar10_resnet18_6 = []
cifar10_resnet18_22 = []
cifar10_resnet18_55 = []
cifar10_resnet18_66 = []
cifar10_resnet18_honest = []
for i in cifar10_resnet18:
    if i[4]=="attacker 1":
        cifar10_resnet18_1.append(i)
    elif i[4]=="attacker 2":
        cifar10_resnet18_2.append(i)
    elif i[4]=="attacker 5":
        cifar10_resnet18_5.append(i)
    elif i[4]=="attacker 6":
        cifar10_resnet18_6.append(i)
    elif i[4]=="attacker 22":
        cifar10_resnet18_22.append(i)
    elif i[4]=="attacker 55":
        cifar10_resnet18_55.append(i)
    elif i[4]=="attacker 66":
        cifar10_resnet18_66.append(i)
    elif i[4]=="honest":
        cifar10_resnet18_honest.append(i)

cifar10_resnet18_1_x = [0]
cifar10_resnet18_1_y = [0.1]
for i in cifar10_resnet18_1:
    cifar10_resnet18_1_x.append(int(i[3])*10)
    cifar10_resnet18_1_y.append(float(i[0]))

# print(cifar10_resnet18_1_x)
# print(cifar10_resnet18_1_y)

cifar10_resnet18_2_x = [0]
cifar10_resnet18_2_y = [0.1]
for i in cifar10_resnet18_2:
    cifar10_resnet18_2_x.append(int(i[3])*10)
    cifar10_resnet18_2_y.append(float(i[0]))

cifar10_resnet18_5_x = [0]
cifar10_resnet18_5_y = [0.1]
for i in cifar10_resnet18_5:
    cifar10_resnet18_5_x.append(int(i[3])*10)
    cifar10_resnet18_5_y.append(float(i[0]))

cifar10_resnet18_6_x = [0]
cifar10_resnet18_6_y = [0.1]
for i in cifar10_resnet18_6:
    cifar10_resnet18_6_x.append(int(i[3])*10)
    cifar10_resnet18_6_y.append(float(i[0]))

cifar10_resnet18_22_x = [0]
cifar10_resnet18_22_y = [0.1]
for i in cifar10_resnet18_22:
    cifar10_resnet18_22_x.append(int(i[3])*10)
    cifar10_resnet18_22_y.append(float(i[0]))

cifar10_resnet18_55_x = [0]
cifar10_resnet18_55_y = [0.1]
for i in cifar10_resnet18_55:
    cifar10_resnet18_55_x.append(int(i[3])*10)
    cifar10_resnet18_55_y.append(float(i[0]))

cifar10_resnet18_66_x = [0]
cifar10_resnet18_66_y = [0.1]
for i in cifar10_resnet18_66:
    cifar10_resnet18_66_x.append(int(i[3])*10)
    cifar10_resnet18_66_y.append(float(i[0]))

cifar10_resnet18_honest_x = [0]
cifar10_resnet18_honest_y = [0.1]
for i in cifar10_resnet18_honest:
    cifar10_resnet18_honest_x.append(int(i[3])*10)
    cifar10_resnet18_honest_y.append(float(i[0]))



font1 = {'family' : 'Times New Roman',
'weight' : f_w,
'size'   : f_s,
}

#开始画图
fig = plt.figure(num=1,figsize=(figsize_x,figsize_y))
# plt.yticks(np.arange(1, y_max, y_em),fontproperties = 'Times New Roman', size = 12)
plt.xticks(np.arange(0, x_max, x_em),fontproperties = 'Times New Roman', size = 12)
# plt.ylim(0, y_max)
plt.xlim(0, x_max)
plt.title('ResNet18')

legend_elements = [
                   Line2D([0], [0], marker='D', color='#1FCA7D', label=a0,
                          markerfacecolor='#1FCA7D', markersize=7), 
                   Line2D([0], [0], marker='^', color='#FBC228', label=a6,
                          markerfacecolor='#FBC228', markersize=7),
                   Line2D([0], [0], marker='o', color='#0000CD', label=a1,
                          markerfacecolor='#0000CD', markersize=7),
                   Line2D([0], [0], marker='s', color='#F54848', label=a2,
                          markerfacecolor='#F54848', markersize=7),
                   Line2D([0], [0], marker='x', color='#01CAFF', label=a5,
                          markerfacecolor='#01CAFF', markersize=7),
                   Line2D([0], [0], color='black', lw=2, linestyle=':', label='Random Guess'),
                   Line2D([0], [0], color='black', lw=2, linestyle='--', label='Shorter Trajectory'),
                   Line2D([0], [0], color='black', lw=2, linestyle='-', label='Longer Trajectory'),]
# Random_Guess_x = [1,10,50,90]
# Random_Guess_y = [0.1,0.1,0.1,0.1]
# plt.plot(Random_Guess_x, Random_Guess_y, marker = '<', color='black', label='Random Guess',linestyle=':',linewidth=2)
plt.axhline(y=0.1, xmin=0, xmax=93, linestyle=':',color = "black")
plt.plot(cifar10_resnet18_1_x, cifar10_resnet18_1_y, marker = 'o', color='#0000CD', label=a1,linestyle='-',linewidth=2)
plt.plot(cifar10_resnet18_2_x, cifar10_resnet18_2_y, marker = 's', color='#F54848', label=a2,linestyle='--',linewidth=2)
plt.plot(cifar10_resnet18_22_x, cifar10_resnet18_22_y, marker = 's', color='#F54848', label=a2,linestyle='-',linewidth=2)
plt.plot(cifar10_resnet18_5_x, cifar10_resnet18_5_y, marker = 'x', color='#01CAFF', label=a5,linestyle='-',linewidth=2)
plt.plot(cifar10_resnet18_55_x, cifar10_resnet18_55_y, marker = 'x', color='#01CAFF', label=a5,linestyle='--',linewidth=2)
plt.plot(cifar10_resnet18_6_x, cifar10_resnet18_6_y, marker = '^', color='#FBC228', label=a6,linestyle='-',linewidth=2)
plt.plot(cifar10_resnet18_66_x, cifar10_resnet18_66_y, marker = '^', color='#FBC228', label=a6,linestyle='--',linewidth=2)
plt.plot(cifar10_resnet18_honest_x, cifar10_resnet18_honest_y, marker = 'D', color='#1FCA7D', label=a0,linestyle='-',linewidth=2)

#plt.plot(sub_axix, test_acys, color='red', label='testing accuracy')

#plt.plot(x_axix, C,  color='#32CD32', label='Normalized Cost ($C\\times p/60$)',linestyle='--',linewidth=3)
#plt.plot(x_axix, T, color='#0000CD', label='Normalized Throughput ($T_2/20$)',linewidth=3)

plt.legend(handles=legend_elements, prop=font1,loc = 'upper center',bbox_to_anchor=(0.78,1)) # 显示图例

plt.xlabel('Synthetic Data Size (SDS)',size='16')
plt.ylabel('Average Accuracy',size='16')
plt.show()

fig.savefig("Documents/IPP/code/pic/cifar10_ResNet18.pdf",dpi=1200)
#endregion











#region cifar10 resnet34
cifar10_resnet34_1 = []
cifar10_resnet34_2 = []
cifar10_resnet34_5 = []
cifar10_resnet34_6 = []
cifar10_resnet34_22 = []
cifar10_resnet34_55 = []
cifar10_resnet34_66 = []
cifar10_resnet34_honest = []
for i in cifar10_resnet34:
    if i[4]=="attacker 1":
        cifar10_resnet34_1.append(i)
    elif i[4]=="attacker 2":
        cifar10_resnet34_2.append(i)
    elif i[4]=="attacker 5":
        cifar10_resnet34_5.append(i)
    elif i[4]=="attacker 6":
        cifar10_resnet34_6.append(i)
    elif i[4]=="attacker 22":
        cifar10_resnet34_22.append(i)
    elif i[4]=="attacker 55":
        cifar10_resnet34_55.append(i)
    elif i[4]=="attacker 66":
        cifar10_resnet34_66.append(i)
    elif i[4]=="honest":
        cifar10_resnet34_honest.append(i)


cifar10_resnet34_1_x = [0]
cifar10_resnet34_1_y = [0.1]
for i in cifar10_resnet34_1:
    cifar10_resnet34_1_x.append(int(i[3])*10)
    cifar10_resnet34_1_y.append(float(i[0]))
                                
cifar10_resnet34_2_x = [0]
cifar10_resnet34_2_y = [0.1]
for i in cifar10_resnet34_2:
    cifar10_resnet34_2_x.append(int(i[3])*10)
    cifar10_resnet34_2_y.append(float(i[0]))

cifar10_resnet34_5_x = [0]
cifar10_resnet34_5_y = [0.1]
for i in cifar10_resnet34_5:
    cifar10_resnet34_5_x.append(int(i[3])*10)
    cifar10_resnet34_5_y.append(float(i[0]))

cifar10_resnet34_6_x = [0]
cifar10_resnet34_6_y = [0.1]
for i in cifar10_resnet34_6:
    cifar10_resnet34_6_x.append(int(i[3])*10)
    cifar10_resnet34_6_y.append(float(i[0]))

cifar10_resnet34_22_x = [0]
cifar10_resnet34_22_y = [0.1]
for i in cifar10_resnet34_22:
    cifar10_resnet34_22_x.append(int(i[3])*10)
    cifar10_resnet34_22_y.append(float(i[0]))

cifar10_resnet34_55_x = [0]
cifar10_resnet34_55_y = [0.1]
for i in cifar10_resnet34_55:
    cifar10_resnet34_55_x.append(int(i[3])*10)
    cifar10_resnet34_55_y.append(float(i[0]))

cifar10_resnet34_66_x = [0]
cifar10_resnet34_66_y = [0.1]
for i in cifar10_resnet34_66:
    cifar10_resnet34_66_x.append(int(i[3])*10)
    cifar10_resnet34_66_y.append(float(i[0]))

cifar10_resnet34_honest_x = [0]
cifar10_resnet34_honest_y = [0.1]
for i in cifar10_resnet34_honest:
    cifar10_resnet34_honest_x.append(int(i[3])*10)
    cifar10_resnet34_honest_y.append(float(i[0]))


font1 = {'family' : 'Times New Roman',
'weight' : f_w,
'size'   : f_s,
}

#开始画图
fig = plt.figure(num=1,figsize=(figsize_x,figsize_y))
# plt.yticks(np.arange(0, x_max, 500),fontproperties = 'Times New Roman', size = 12)
plt.xticks(np.arange(0, x_max, x_em),fontproperties = 'Times New Roman', size = 12)
# plt.ylim(0, 3500)
plt.xlim(0, x_max)
plt.title('ResNet34')

legend_elements = [
                   Line2D([0], [0], marker='D', color='#1FCA7D', label=a0,
                          markerfacecolor='#1FCA7D', markersize=7), 
                   Line2D([0], [0], marker='^', color='#FBC228', label=a6,
                          markerfacecolor='#FBC228', markersize=7),
                   Line2D([0], [0], marker='o', color='#0000CD', label=a1,
                          markerfacecolor='#0000CD', markersize=7),
                   Line2D([0], [0], marker='s', color='#F54848', label=a2,
                          markerfacecolor='#F54848', markersize=7),
                   Line2D([0], [0], marker='x', color='#01CAFF', label=a5,
                          markerfacecolor='#01CAFF', markersize=7),
                   Line2D([0], [0], color='black', lw=2, linestyle=':', label='Random Guess'),
                   Line2D([0], [0], color='black', lw=2, linestyle='--', label='Shorter Trajectory'),
                   Line2D([0], [0], color='black', lw=2, linestyle='-', label='Longer Trajectory'),]
# Random_Guess_x = [1,10,50,90]
# Random_Guess_y = [0.1,0.1,0.1,0.1]

# plt.plot(Random_Guess_x, Random_Guess_y, marker = '<', color='black', label='Random Guess',linestyle=':',linewidth=2)
plt.axhline(y=0.1, xmin=0, xmax=93, linestyle=':',color = "black")
plt.plot(cifar10_resnet34_1_x, cifar10_resnet34_1_y, marker = 'o', color='#0000CD', label=a1,linestyle='-',linewidth=2)
plt.plot(cifar10_resnet34_2_x, cifar10_resnet34_2_y, marker = 's', color='#F54848', label=a2,linestyle='--',linewidth=2)
plt.plot(cifar10_resnet34_22_x, cifar10_resnet34_22_y, marker = 's', color='#F54848', label=a2,linestyle='-',linewidth=2)
plt.plot(cifar10_resnet34_5_x, cifar10_resnet34_5_y, marker = 'x', color='#01CAFF', label=a5,linestyle='-',linewidth=2)
plt.plot(cifar10_resnet34_55_x, cifar10_resnet34_55_y, marker = 'x', color='#01CAFF', label=a5,linestyle='--',linewidth=2)
plt.plot(cifar10_resnet34_6_x, cifar10_resnet34_6_y, marker = '^', color='#FBC228', label=a6,linestyle='-',linewidth=2)
plt.plot(cifar10_resnet34_66_x, cifar10_resnet34_66_y, marker = '^', color='#FBC228', label=a6,linestyle='--',linewidth=2)
plt.plot(cifar10_resnet34_honest_x, cifar10_resnet34_honest_y, marker = 'D', color='#1FCA7D', label=a0,linestyle='-',linewidth=2)

#plt.plot(sub_axix, test_acys, color='red', label='testing accuracy')

#plt.plot(x_axix, C,  color='#32CD32', label='Normalized Cost ($C\\times p/60$)',linestyle='--',linewidth=3)
#plt.plot(x_axix, T, color='#0000CD', label='Normalized Throughput ($T_2/20$)',linewidth=3)

plt.legend(handles=legend_elements, prop=font1,loc = 'upper center',bbox_to_anchor=(0.78,1)) # 显示图例

plt.xlabel('Synthetic Data Size (SDS)',size='16')
plt.ylabel('Average Accuracy',size='16')
plt.show()
fig.savefig("Documents/IPP/code/pic/cifar10_ResNet34.pdf",dpi=1200)
#endregion











#region cifar10 convnet
cifar10_convnet_1 = []
cifar10_convnet_2 = []
cifar10_convnet_5 = []
cifar10_convnet_6 = []
cifar10_convnet_22 = []
cifar10_convnet_55 = []
cifar10_convnet_66 = []
cifar10_convnet_honest = []
for i in cifar10_convnet:
    if i[4]=="attacker 1":
        cifar10_convnet_1.append(i)
    elif i[4]=="attacker 2":
        cifar10_convnet_2.append(i)
    elif i[4]=="attacker 5":
        cifar10_convnet_5.append(i)
    elif i[4]=="attacker 6":
        cifar10_convnet_6.append(i)
    elif i[4]=="attacker 22":
        cifar10_convnet_22.append(i)
    elif i[4]=="attacker 55":
        cifar10_convnet_55.append(i)
    elif i[4]=="attacker 66":
        cifar10_convnet_66.append(i)
    elif i[4]=="honest":
        cifar10_convnet_honest.append(i)

cifar10_convnet_1_x = []
cifar10_convnet_1_y = []
for i in cifar10_convnet_1:
    cifar10_convnet_1_x.append(int(i[3])*10)
    cifar10_convnet_1_y.append(float(i[0]))
                                
cifar10_convnet_2_x = []
cifar10_convnet_2_y = []
for i in cifar10_convnet_2:
    cifar10_convnet_2_x.append(int(i[3])*10)
    cifar10_convnet_2_y.append(float(i[0]))

# print(cifar10_convnet_2_x)
# print(cifar10_convnet_2_y)
# hj

cifar10_convnet_5_x = []
cifar10_convnet_5_y = []
for i in cifar10_convnet_5:
    cifar10_convnet_5_x.append(int(i[3])*10)
    cifar10_convnet_5_y.append(float(i[0]))

cifar10_convnet_6_x = []
cifar10_convnet_6_y = []
for i in cifar10_convnet_6:
    cifar10_convnet_6_x.append(int(i[3])*10)
    cifar10_convnet_6_y.append(float(i[0]))

cifar10_convnet_22_x = []
cifar10_convnet_22_y = []
for i in cifar10_convnet_22:
    cifar10_convnet_22_x.append(int(i[3])*10)
    cifar10_convnet_22_y.append(float(i[0]))

cifar10_convnet_55_x = []
cifar10_convnet_55_y = []
for i in cifar10_convnet_55:
    cifar10_convnet_55_x.append(int(i[3])*10)
    cifar10_convnet_55_y.append(float(i[0]))

cifar10_convnet_66_x = []
cifar10_convnet_66_y = []
for i in cifar10_convnet_66:
    cifar10_convnet_66_x.append(int(i[3])*10)
    cifar10_convnet_66_y.append(float(i[0]))

cifar10_convnet_honest_x = []
cifar10_convnet_honest_y = []
for i in cifar10_convnet_honest:
    cifar10_convnet_honest_x.append(int(i[3])*10)
    cifar10_convnet_honest_y.append(float(i[0]))


font1 = {'family' : 'Times New Roman',
'weight' : f_w,
'size'   : f_s,
}

#开始画图
fig = plt.figure(num=1,figsize=(figsize_x,figsize_y))
# plt.yticks(np.arange(0, 3500, 500),fontproperties = 'Times New Roman', size = 12)
plt.xticks(np.arange(0, x_max, x_em),fontproperties = 'Times New Roman', size = 12)
# plt.ylim(0, 3500)
plt.xlim(0, x_max)
plt.title('ConvNet')

legend_elements = [
                   Line2D([0], [0], marker='D', color='#1FCA7D', label=a0,
                          markerfacecolor='#1FCA7D', markersize=7), 
                   Line2D([0], [0], marker='^', color='#FBC228', label=a6,
                          markerfacecolor='#FBC228', markersize=7),
                   Line2D([0], [0], marker='o', color='#0000CD', label=a1,
                          markerfacecolor='#0000CD', markersize=7),
                   Line2D([0], [0], marker='s', color='#F54848', label=a2,
                          markerfacecolor='#F54848', markersize=7),
                   Line2D([0], [0], marker='x', color='#01CAFF', label=a5,
                          markerfacecolor='#01CAFF', markersize=7),
                   Line2D([0], [0], color='black', lw=2, linestyle=':', label='Random Guess'),
                   Line2D([0], [0], color='black', lw=2, linestyle='--', label='Shorter Trajectory'),
                   Line2D([0], [0], color='black', lw=2, linestyle='-', label='Longer Trajectory'),]
# Random_Guess_x = [1,10,50,90]
# Random_Guess_y = [0.1,0.1,0.1,0.1]
fig = plt.figure(num=1,figsize=(10,5))
# plt.plot(Random_Guess_x, Random_Guess_y, marker = '<', color='black', label='Random Guess',linestyle=':',linewidth=2)
plt.axhline(y=0.1, xmin=0, xmax=93, linestyle=':',color = "black")
plt.plot(cifar10_convnet_1_x, cifar10_convnet_1_y, marker = 'o', color='#0000CD', label=a1,linestyle='-',linewidth=2)
plt.plot(cifar10_convnet_2_x, cifar10_convnet_2_y, marker = 's', color='#F54848', label=a2,linestyle='--',linewidth=2)
plt.plot(cifar10_convnet_22_x, cifar10_convnet_22_y, marker = 's', color='#F54848', label=a2,linestyle='-',linewidth=2)
plt.plot(cifar10_convnet_5_x, cifar10_convnet_5_y, marker = 'x', color='#01CAFF', label=a5,linestyle='-',linewidth=2)
plt.plot(cifar10_convnet_55_x, cifar10_convnet_55_y, marker = 'x', color='#01CAFF', label=a5,linestyle='--',linewidth=2)
plt.plot(cifar10_convnet_6_x, cifar10_convnet_6_y, marker = '^', color='#FBC228', label=a6,linestyle='-',linewidth=2)
plt.plot(cifar10_convnet_66_x, cifar10_convnet_66_y, marker = '^', color='#FBC228', label=a6,linestyle='--',linewidth=2)
plt.plot(cifar10_convnet_honest_x, cifar10_convnet_honest_y, marker = 'D', color='#1FCA7D', label=a0,linestyle='-',linewidth=2)

#plt.plot(sub_axix, test_acys, color='red', label='testing accuracy')

#plt.plot(x_axix, C,  color='#32CD32', label='Normalized Cost ($C\\times p/60$)',linestyle='--',linewidth=3)
#plt.plot(x_axix, T, color='#0000CD', label='Normalized Throughput ($T_2/20$)',linewidth=3)

plt.legend(handles=legend_elements, prop=font1,loc = 'upper center',bbox_to_anchor=(0.78,1)) # 显示图例

plt.xlabel('Synthetic Data Size (SDS)',size='16')
plt.ylabel('Average Accuracy',size='16')
plt.show()
fig.savefig("Documents/IPP/code/pic/cifar10_CNN.pdf",dpi=1200)
#endregion









# region cifar10 convnetd2
cifar10_convnetd2_1 = []
cifar10_convnetd2_2 = []
cifar10_convnetd2_5 = []
cifar10_convnetd2_6 = []
cifar10_convnetd2_22 = []
cifar10_convnetd2_55 = []
cifar10_convnetd2_66 = []
cifar10_convnetd2_honest = []
for i in cifar10_convnetd2:
    if i[4]=="attacker 1":
        cifar10_convnetd2_1.append(i)
    elif i[4]=="attacker 2":
        cifar10_convnetd2_2.append(i)
    elif i[4]=="attacker 5":
        cifar10_convnetd2_5.append(i)
    elif i[4]=="attacker 6":
        cifar10_convnetd2_6.append(i)
    elif i[4]=="attacker 22":
        cifar10_convnetd2_22.append(i)
    elif i[4]=="attacker 55":
        cifar10_convnetd2_55.append(i)
    elif i[4]=="attacker 66":
        cifar10_convnetd2_66.append(i)
    elif i[4]=="honest":
        cifar10_convnetd2_honest.append(i)


cifar10_convnetd2_1_x = [0]
cifar10_convnetd2_1_y = [0.1]
for i in cifar10_convnetd2_1:
    cifar10_convnetd2_1_x.append(int(i[3])*10)
    cifar10_convnetd2_1_y.append(float(i[0]))
                                
cifar10_convnetd2_2_x = [0]
cifar10_convnetd2_2_y = [0.1]
for i in cifar10_convnetd2_2:
    cifar10_convnetd2_2_x.append(int(i[3])*10)
    cifar10_convnetd2_2_y.append(float(i[0]))

cifar10_convnetd2_5_x = [0]
cifar10_convnetd2_5_y = [0.1]
for i in cifar10_convnetd2_5:
    cifar10_convnetd2_5_x.append(int(i[3])*10)
    cifar10_convnetd2_5_y.append(float(i[0]))

cifar10_convnetd2_6_x = [0]
cifar10_convnetd2_6_y = [0.1]
for i in cifar10_convnetd2_6:
    cifar10_convnetd2_6_x.append(int(i[3])*10)
    cifar10_convnetd2_6_y.append(float(i[0]))

cifar10_convnetd2_22_x = [0]
cifar10_convnetd2_22_y = [0.1]
for i in cifar10_convnetd2_22:
    cifar10_convnetd2_22_x.append(int(i[3])*10)
    cifar10_convnetd2_22_y.append(float(i[0]))

cifar10_convnetd2_55_x = [0]
cifar10_convnetd2_55_y = [0.1]
for i in cifar10_convnetd2_55:
    cifar10_convnetd2_55_x.append(int(i[3])*10)
    cifar10_convnetd2_55_y.append(float(i[0]))

cifar10_convnetd2_66_x = [0]
cifar10_convnetd2_66_y = [0.1]
for i in cifar10_convnetd2_66:
    cifar10_convnetd2_66_x.append(int(i[3])*10)
    cifar10_convnetd2_66_y.append(float(i[0]))

cifar10_convnetd2_honest_x = [0]
cifar10_convnetd2_honest_y = [0.1]
for i in cifar10_convnetd2_honest:
    cifar10_convnetd2_honest_x.append(int(i[3])*10)
    cifar10_convnetd2_honest_y.append(float(i[0]))


font1 = {'family' : 'Times New Roman',
'weight' : f_w,
'size'   : f_s,
}

#开始画图
fig = plt.figure(num=1,figsize=(figsize_x,figsize_y))
# plt.yticks(np.arange(0, 3500, 500),fontproperties = 'Times New Roman', size = 12)
plt.xticks(np.arange(0, x_max, x_em),fontproperties = 'Times New Roman', size = 12)
# plt.ylim(0, 3500)
plt.xlim(0, x_max)
plt.title('ConvNetD2')

legend_elements = [
                   Line2D([0], [0], marker='D', color='#1FCA7D', label=a0,
                          markerfacecolor='#1FCA7D', markersize=7), 
                   Line2D([0], [0], marker='^', color='#FBC228', label=a6,
                          markerfacecolor='#FBC228', markersize=7),
                   Line2D([0], [0], marker='o', color='#0000CD', label=a1,
                          markerfacecolor='#0000CD', markersize=7),
                   Line2D([0], [0], marker='s', color='#F54848', label=a2,
                          markerfacecolor='#F54848', markersize=7),
                   Line2D([0], [0], marker='x', color='#01CAFF', label=a5,
                          markerfacecolor='#01CAFF', markersize=7),
                   Line2D([0], [0], color='black', lw=2, linestyle=':', label='Random Guess'),
                   Line2D([0], [0], color='black', lw=2, linestyle='--', label='Shorter Trajectory'),
                   Line2D([0], [0], color='black', lw=2, linestyle='-', label='Longer Trajectory'),]
# Random_Guess_x = [1,10,50,90]
# Random_Guess_y = [0.1,0.1,0.1,0.1]

# plt.plot(Random_Guess_x, Random_Guess_y, marker = '<', color='black', label='Random Guess',linestyle=':',linewidth=2)
plt.axhline(y=0.1, xmin=0, xmax=93, linestyle=':',color = "black")
plt.plot(cifar10_convnetd2_1_x, cifar10_convnetd2_1_y, marker = 'o', color='#0000CD', label=a1,linestyle='-',linewidth=2)
plt.plot(cifar10_convnetd2_2_x, cifar10_convnetd2_2_y, marker = 's', color='#F54848', label=a2,linestyle='--',linewidth=2)
plt.plot(cifar10_convnetd2_22_x, cifar10_convnetd2_22_y, marker = 's', color='#F54848', label=a2,linestyle='-',linewidth=2)
plt.plot(cifar10_convnetd2_5_x, cifar10_convnetd2_5_y, marker = 'x', color='#01CAFF', label=a5,linestyle='-',linewidth=2)
plt.plot(cifar10_convnetd2_55_x, cifar10_convnetd2_55_y, marker = 'x', color='#01CAFF', label=a5,linestyle='--',linewidth=2)
plt.plot(cifar10_convnetd2_6_x, cifar10_convnetd2_6_y, marker = '^', color='#FBC228', label=a6,linestyle='-',linewidth=2)
plt.plot(cifar10_convnetd2_66_x, cifar10_convnetd2_66_y, marker = '^', color='#FBC228', label=a6,linestyle='--',linewidth=2)
plt.plot(cifar10_convnetd2_honest_x, cifar10_convnetd2_honest_y, marker = 'D', color='#1FCA7D', label=a0,linestyle='-',linewidth=2)

#plt.plot(sub_axix, test_acys, color='red', label='testing accuracy')

#plt.plot(x_axix, C,  color='#32CD32', label='Normalized Cost ($C\\times p/60$)',linestyle='--',linewidth=3)
#plt.plot(x_axix, T, color='#0000CD', label='Normalized Throughput ($T_2/20$)',linewidth=3)

plt.legend(handles=legend_elements, prop=font1,loc = 'upper center',bbox_to_anchor=(0.78,1)) # 显示图例

plt.xlabel('Synthetic Data Size (SDS)',size='16')
plt.ylabel('Average Accuracy',size='16')
plt.show()
fig.savefig("Documents/IPP/code/pic/cifar10_CNN2D.pdf",dpi=1200)
#endregion
       



# region cifar10 convnetd4
cifar10_convnetd4_1 = []
cifar10_convnetd4_2 = []
cifar10_convnetd4_5 = []
cifar10_convnetd4_6 = []
cifar10_convnetd4_22 = []
cifar10_convnetd4_55 = []
cifar10_convnetd4_66 = []
cifar10_convnetd4_honest = []
for i in cifar10_convnetd4:
    if i[4]=="attacker 1":
        cifar10_convnetd4_1.append(i)
    elif i[4]=="attacker 2":
        cifar10_convnetd4_2.append(i)
    elif i[4]=="attacker 5":
        cifar10_convnetd4_5.append(i)
    elif i[4]=="attacker 6":
        cifar10_convnetd4_6.append(i)
    elif i[4]=="attacker 22":
        cifar10_convnetd4_22.append(i)
    elif i[4]=="attacker 55":
        cifar10_convnetd4_55.append(i)
    elif i[4]=="attacker 66":
        cifar10_convnetd4_66.append(i)
    elif i[4]=="honest":
        cifar10_convnetd4_honest.append(i)


cifar10_convnetd4_1_x = [0]
cifar10_convnetd4_1_y = [0.1]
for i in cifar10_convnetd4_1:
    cifar10_convnetd4_1_x.append(int(i[3])*10)
    cifar10_convnetd4_1_y.append(float(i[0]))
                                
cifar10_convnetd4_2_x = [0]
cifar10_convnetd4_2_y = [0.1]
for i in cifar10_convnetd4_2:
    cifar10_convnetd4_2_x.append(int(i[3])*10)
    cifar10_convnetd4_2_y.append(float(i[0]))

cifar10_convnetd4_5_x = [0]
cifar10_convnetd4_5_y = [0.1]
for i in cifar10_convnetd4_5:
    cifar10_convnetd4_5_x.append(int(i[3])*10)
    cifar10_convnetd4_5_y.append(float(i[0]))

cifar10_convnetd4_6_x = [0]
cifar10_convnetd4_6_y = [0.1]
for i in cifar10_convnetd4_6:
    cifar10_convnetd4_6_x.append(int(i[3])*10)
    cifar10_convnetd4_6_y.append(float(i[0]))

cifar10_convnetd4_22_x = [0]
cifar10_convnetd4_22_y = [0.1]
for i in cifar10_convnetd4_22:
    cifar10_convnetd4_22_x.append(int(i[3])*10)
    cifar10_convnetd4_22_y.append(float(i[0]))

cifar10_convnetd4_55_x = [0]
cifar10_convnetd4_55_y = [0.1]
for i in cifar10_convnetd4_55:
    cifar10_convnetd4_55_x.append(int(i[3])*10)
    cifar10_convnetd4_55_y.append(float(i[0]))

cifar10_convnetd4_66_x = [0]
cifar10_convnetd4_66_y = [0.1]
for i in cifar10_convnetd4_66:
    cifar10_convnetd4_66_x.append(int(i[3])*10)
    cifar10_convnetd4_66_y.append(float(i[0]))

cifar10_convnetd4_honest_x = [0]
cifar10_convnetd4_honest_y = [0.1]
for i in cifar10_convnetd4_honest:
    cifar10_convnetd4_honest_x.append(int(i[3])*10)
    cifar10_convnetd4_honest_y.append(float(i[0]))


font1 = {'family' : 'Times New Roman',
'weight' : f_w,
'size'   : f_s,
}

#开始画图
fig = plt.figure(num=1,figsize=(figsize_x,figsize_y))
# plt.yticks(np.arange(0, 3500, 500),fontproperties = 'Times New Roman', size = 12)
plt.xticks(np.arange(0, x_max, x_em),fontproperties = 'Times New Roman', size = 12)
# plt.ylim(0, 3500)
plt.xlim(0, x_max)
plt.title('ConNetD4')

legend_elements = [
                   Line2D([0], [0], marker='D', color='#1FCA7D', label=a0,
                          markerfacecolor='#1FCA7D', markersize=7), 
                   Line2D([0], [0], marker='^', color='#FBC228', label=a6,
                          markerfacecolor='#FBC228', markersize=7),
                   Line2D([0], [0], marker='o', color='#0000CD', label=a1,
                          markerfacecolor='#0000CD', markersize=7),
                   Line2D([0], [0], marker='s', color='#F54848', label=a2,
                          markerfacecolor='#F54848', markersize=7),
                   Line2D([0], [0], marker='x', color='#01CAFF', label=a5,
                          markerfacecolor='#01CAFF', markersize=7),
                   Line2D([0], [0], color='black', lw=2, linestyle=':', label='Random Guess'),
                   Line2D([0], [0], color='black', lw=2, linestyle='--', label='Shorter Trajectory'),
                   Line2D([0], [0], color='black', lw=2, linestyle='-', label='Longer Trajectory'),]
# Random_Guess_x = [1,10,50,90]
# Random_Guess_y = [0.1,0.1,0.1,0.1]

# plt.plot(Random_Guess_x, Random_Guess_y, marker = '<', color='black', label='Random Guess',linestyle=':',linewidth=2)
plt.axhline(y=0.1, xmin=0, xmax=93, linestyle=':',color = "black")
plt.plot(cifar10_convnetd4_1_x, cifar10_convnetd4_1_y, marker = 'o', color='#0000CD', label=a1,linestyle='-',linewidth=2)
plt.plot(cifar10_convnetd4_2_x, cifar10_convnetd4_2_y, marker = 's', color='#F54848', label=a2,linestyle='--',linewidth=2)
plt.plot(cifar10_convnetd4_22_x, cifar10_convnetd4_22_y, marker = 's', color='#F54848', label=a2,linestyle='-',linewidth=2)
plt.plot(cifar10_convnetd4_5_x, cifar10_convnetd4_5_y, marker = 'x', color='#01CAFF', label=a5,linestyle='-',linewidth=2)
plt.plot(cifar10_convnetd4_55_x, cifar10_convnetd4_55_y, marker = 'x', color='#01CAFF', label=a5,linestyle='--',linewidth=2)
plt.plot(cifar10_convnetd4_6_x, cifar10_convnetd4_6_y, marker = '^', color='#FBC228', label=a6,linestyle='-',linewidth=2)
plt.plot(cifar10_convnetd4_66_x, cifar10_convnetd4_66_y, marker = '^', color='#FBC228', label=a6,linestyle='--',linewidth=2)
plt.plot(cifar10_convnetd4_honest_x, cifar10_convnetd4_honest_y, marker = 'D', color='#1FCA7D', label=a0,linestyle='-',linewidth=2)

#plt.plot(sub_axix, test_acys, color='red', label='testing accuracy')

#plt.plot(x_axix, C,  color='#32CD32', label='Normalized Cost ($C\\times p/60$)',linestyle='--',linewidth=3)
#plt.plot(x_axix, T, color='#0000CD', label='Normalized Throughput ($T_2/20$)',linewidth=3)

plt.legend(handles=legend_elements, prop=font1,loc = 'upper center',bbox_to_anchor=(0.78,1)) # 显示图例

plt.xlabel('Synthetic Data Size (SDS)',size='16')
plt.ylabel('Average Accuracy',size='16')
plt.show()
fig.savefig("Documents/IPP/code/pic/cifar10_CNN4D.pdf",dpi=1200)
#endregion
  















# # region cifar100 resnet18
# cifar100_resnet18_1 = []
# cifar100_resnet18_2 = []
# cifar100_resnet18_5 = []
# cifar100_resnet18_6 = []
# cifar100_resnet18_22 = []
# cifar100_resnet18_55 = []
# cifar100_resnet18_66 = []
# cifar100_resnet18_honest = []
# for i in cifar100_resnet18:
#     if i[4]=="attacker 1":
#         cifar100_resnet18_1.append(i)
#     elif i[4]=="attacker 2":
#         cifar100_resnet18_2.append(i)
#     elif i[4]=="attacker 5":
#         cifar100_resnet18_5.append(i)
#     elif i[4]=="attacker 6":
#         cifar100_resnet18_6.append(i)
#     elif i[4]=="attacker 22":
#         cifar100_resnet18_22.append(i)
#     elif i[4]=="attacker 55":
#         cifar100_resnet18_55.append(i)
#     elif i[4]=="attacker 66":
#         cifar100_resnet18_66.append(i)
#     elif i[4]=="honest":
#         cifar100_resnet18_honest.append(i)
# #endregion

# cifar100_resnet18_1_x = []
# cifar100_resnet18_1_y = []
# for i in cifar100_resnet18_1:
#     cifar100_resnet18_1_x.append(int(i[3]))
#     cifar100_resnet18_1_y.append(float(i[0]))

# # print(cifar100_resnet18_1_x)
# # print(cifar100_resnet18_1_y)

# cifar100_resnet18_2_x = []
# cifar100_resnet18_2_y = []
# for i in cifar100_resnet18_2:
#     cifar100_resnet18_2_x.append(int(i[3]))
#     cifar100_resnet18_2_y.append(float(i[0]))

# cifar100_resnet18_5_x = []
# cifar100_resnet18_5_y = []
# for i in cifar100_resnet18_5:
#     cifar100_resnet18_5_x.append(int(i[3]))
#     cifar100_resnet18_5_y.append(float(i[0]))

# cifar100_resnet18_6_x = []
# cifar100_resnet18_6_y = []
# for i in cifar100_resnet18_6:
#     cifar100_resnet18_6_x.append(int(i[3]))
#     cifar100_resnet18_6_y.append(float(i[0]))

# cifar100_resnet18_22_x = []
# cifar100_resnet18_22_y = []
# for i in cifar100_resnet18_22:
#     cifar100_resnet18_22_x.append(int(i[3]))
#     cifar100_resnet18_22_y.append(float(i[0]))

# cifar100_resnet18_55_x = []
# cifar100_resnet18_55_y = []
# for i in cifar100_resnet18_55:
#     cifar100_resnet18_55_x.append(int(i[3]))
#     cifar100_resnet18_55_y.append(float(i[0]))

# cifar100_resnet18_66_x = []
# cifar100_resnet18_66_y = []
# for i in cifar100_resnet18_66:
#     cifar100_resnet18_66_x.append(int(i[3]))
#     cifar100_resnet18_66_y.append(float(i[0]))

# cifar100_resnet18_honest_x = []
# cifar100_resnet18_honest_y = []
# for i in cifar100_resnet18_honest:
#     cifar100_resnet18_honest_x.append(int(i[3]))
#     cifar100_resnet18_honest_y.append(float(i[0]))



# font1 = {'family' : 'Times New Roman',
# 'weight' : 'normal',
# 'size'   : 14,
# }

# #开始画图
# # plt.yticks(np.arange(1, x_max, 10),fontproperties = 'Times New Roman', size = 12)
# plt.xticks(np.arange(0, x_max, x_em),fontproperties = 'Times New Roman', size = 12)
# # plt.ylim(0, 3500)
# plt.xlim(0, x_max)
# plt.title('ResNet18')

# legend_elements = [Line2D([0], [0], marker='o', color='#32CD32', label=a1,
#                           markerfacecolor='#32CD32', markersize=7),
#                    Line2D([0], [0], marker='s', color='#0000CD', label=a2,
#                           markerfacecolor='#0000CD', markersize=7),
#                    Line2D([0], [0], marker='x', color='red', label=a5,
#                           markerfacecolor='red', markersize=7),
#                    Line2D([0], [0], marker='^', color='#FFCC33', label=a6,
#                           markerfacecolor='#FFCC33', markersize=7),
#                    Line2D([0], [0], marker='D', color='#1A1A1A', label=a0,
#                           markerfacecolor='#1A1A1A', markersize=7),
#                    Line2D([0], [0], color='black', lw=2, linestyle=':', label='longer'),
#                    Line2D([0], [0], color='black', lw=2, linestyle='--', label='shorter'),
#                    Line2D([0], [0], color='black', lw=2, linestyle='-', label='normal'),]

# plt.plot(cifar100_resnet18_1_x, cifar100_resnet18_1_y, marker = 'o', color='#32CD32', label=a1,linestyle='-',linewidth=2)
# plt.plot(cifar100_resnet18_2_x, cifar100_resnet18_2_y, marker = 's', color='#0000CD', label=a2,linestyle='-',linewidth=2)
# plt.plot(cifar100_resnet18_22_x, cifar100_resnet18_22_y, marker = 's', color='#0000CD', label=a2,linestyle=':',linewidth=2)
# plt.plot(cifar100_resnet18_5_x, cifar100_resnet18_5_y, marker = 'x', color='red', label=a5,linestyle='-',linewidth=2)
# plt.plot(cifar100_resnet18_55_x, cifar100_resnet18_55_y, marker = 'x', color='red', label=a5,linestyle='--',linewidth=2)
# plt.plot(cifar100_resnet18_6_x, cifar100_resnet18_6_y, marker = '^', color='#FFCC33', label=a6,linestyle='-',linewidth=2)
# plt.plot(cifar100_resnet18_66_x, cifar100_resnet18_66_y, marker = '^', color='#FFCC33', label=a6,linestyle='--',linewidth=2)
# plt.plot(cifar100_resnet18_honest_x, cifar100_resnet18_honest_y, marker = 'D', color='#1A1A1A', label=a0,linestyle='-',linewidth=2)

# #plt.plot(sub_axix, test_acys, color='red', label='testing accuracy')

# #plt.plot(x_axix, C,  color='#32CD32', label='Normalized Cost ($C\\times p/60$)',linestyle='--',linewidth=3)
# #plt.plot(x_axix, T, color='#0000CD', label='Normalized Throughput ($T_2/20$)',linewidth=3)

# plt.legend(handles=legend_elements, prop=font1,loc = 'upper center',bbox_to_anchor=(0.78,1)) # 显示图例

# plt.xlabel('Number of Images per Class (IPC)',size='16')
# plt.ylabel('Average Accuracy',size='16')
# plt.show()
# # fig.savefig("result2.pdf",dpi=1200)
# #endregion




# #region cifar100 resnet34
# cifar100_resnet34_1 = []
# cifar100_resnet34_2 = []
# cifar100_resnet34_5 = []
# cifar100_resnet34_6 = []
# cifar100_resnet34_22 = []
# cifar100_resnet34_55 = []
# cifar100_resnet34_66 = []
# cifar100_resnet34_honest = []
# for i in cifar100_resnet34:
#     if i[4]=="attacker 1":
#         cifar100_resnet34_1.append(i)
#     elif i[4]=="attacker 2":
#         cifar100_resnet34_2.append(i)
#     elif i[4]=="attacker 5":
#         cifar100_resnet34_5.append(i)
#     elif i[4]=="attacker 6":
#         cifar100_resnet34_6.append(i)
#     elif i[4]=="attacker 22":
#         cifar100_resnet34_22.append(i)
#     elif i[4]=="attacker 55":
#         cifar100_resnet34_55.append(i)
#     elif i[4]=="attacker 66":
#         cifar100_resnet34_66.append(i)
#     elif i[4]=="honest":
#         cifar100_resnet34_honest.append(i)
# # endregion

# cifar100_resnet34_1_x = []
# cifar100_resnet34_1_y = []
# for i in cifar100_resnet34_1:
#     cifar100_resnet34_1_x.append(int(i[3]))
#     cifar100_resnet34_1_y.append(float(i[0]))
                                
# cifar100_resnet34_2_x = []
# cifar100_resnet34_2_y = []
# for i in cifar100_resnet34_2:
#     cifar100_resnet34_2_x.append(int(i[3]))
#     cifar100_resnet34_2_y.append(float(i[0]))

# cifar100_resnet34_5_x = []
# cifar100_resnet34_5_y = []
# for i in cifar100_resnet34_5:
#     cifar100_resnet34_5_x.append(int(i[3]))
#     cifar100_resnet34_5_y.append(float(i[0]))

# cifar100_resnet34_6_x = []
# cifar100_resnet34_6_y = []
# for i in cifar100_resnet34_6:
#     cifar100_resnet34_6_x.append(int(i[3]))
#     cifar100_resnet34_6_y.append(float(i[0]))

# cifar100_resnet34_22_x = []
# cifar100_resnet34_22_y = []
# for i in cifar100_resnet34_22:
#     cifar100_resnet34_22_x.append(int(i[3]))
#     cifar100_resnet34_22_y.append(float(i[0]))

# cifar100_resnet34_55_x = []
# cifar100_resnet34_55_y = []
# for i in cifar100_resnet34_55:
#     cifar100_resnet34_55_x.append(int(i[3]))
#     cifar100_resnet34_55_y.append(float(i[0]))

# cifar100_resnet34_66_x = []
# cifar100_resnet34_66_y = []
# for i in cifar100_resnet34_66:
#     cifar100_resnet34_66_x.append(int(i[3]))
#     cifar100_resnet34_66_y.append(float(i[0]))

# cifar100_resnet34_honest_x = []
# cifar100_resnet34_honest_y = []
# for i in cifar100_resnet34_honest:
#     cifar100_resnet34_honest_x.append(int(i[3]))
#     cifar100_resnet34_honest_y.append(float(i[0]))


# font1 = {'family' : 'Times New Roman',
# 'weight' : 'normal',
# 'size'   : 14,
# }

# #开始画图
# # plt.yticks(np.arange(0, x_max, 500),fontproperties = 'Times New Roman', size = 12)
# plt.xticks(np.arange(0, x_max, x_em),fontproperties = 'Times New Roman', size = 12)
# # plt.ylim(0, 3500)
# plt.xlim(0, x_max)
# plt.title('ResNet34')

# legend_elements = [Line2D([0], [0], marker='o', color='#32CD32', label=a1,
#                           markerfacecolor='#32CD32', markersize=7),
#                    Line2D([0], [0], marker='s', color='#0000CD', label=a2,
#                           markerfacecolor='#0000CD', markersize=7),
#                    Line2D([0], [0], marker='x', color='red', label=a5,
#                           markerfacecolor='red', markersize=7),
#                    Line2D([0], [0], marker='^', color='#FFCC33', label=a6,
#                           markerfacecolor='#FFCC33', markersize=7),
#                    Line2D([0], [0], marker='D', color='#1A1A1A', label=a0,
#                           markerfacecolor='#1A1A1A', markersize=7),
#                    Line2D([0], [0], color='black', lw=2, linestyle=':', label='longer'),
#                    Line2D([0], [0], color='black', lw=2, linestyle='--', label='shorter'),
#                    Line2D([0], [0], color='black', lw=2, linestyle='-', label='normal'),]

# plt.plot(cifar100_resnet34_1_x, cifar100_resnet34_1_y, marker = 'o', color='#32CD32', label=a1,linestyle='-',linewidth=2)
# plt.plot(cifar100_resnet34_2_x, cifar100_resnet34_2_y, marker = 's', color='#0000CD', label=a2,linestyle='-',linewidth=2)
# plt.plot(cifar100_resnet34_22_x, cifar100_resnet34_22_y, marker = 's', color='#0000CD', label=a2,linestyle=':',linewidth=2)
# plt.plot(cifar100_resnet34_5_x, cifar100_resnet34_5_y, marker = 'x', color='red', label=a5,linestyle='-',linewidth=2)
# plt.plot(cifar100_resnet34_55_x, cifar100_resnet34_55_y, marker = 'x', color='red', label=a5,linestyle='--',linewidth=2)
# plt.plot(cifar100_resnet34_6_x, cifar100_resnet34_6_y, marker = '^', color='#FFCC33', label=a6,linestyle='-',linewidth=2)
# plt.plot(cifar100_resnet34_66_x, cifar100_resnet34_66_y, marker = '^', color='#FFCC33', label=a6,linestyle='--',linewidth=2)
# plt.plot(cifar100_resnet34_honest_x, cifar100_resnet34_honest_y, marker = 'D', color='#1A1A1A', label=a0,linestyle='-',linewidth=2)

# #plt.plot(sub_axix, test_acys, color='red', label='testing accuracy')

# #plt.plot(x_axix, C,  color='#32CD32', label='Normalized Cost ($C\\times p/60$)',linestyle='--',linewidth=3)
# #plt.plot(x_axix, T, color='#0000CD', label='Normalized Throughput ($T_2/20$)',linewidth=3)
# plt.legend(handles=legend_elements, prop=font1,loc = 'upper center',bbox_to_anchor=(0.78,1)) # 显示图例

# plt.xlabel('Number of Images per Class (IPC)',size='16')
# plt.ylabel('Average Accuracy',size='16')
# plt.show()
# # fig.savefig("result2.pdf",dpi=1200)

# # endregion





# #region cifar100 convnet
# cifar100_convnet_1 = []
# cifar100_convnet_2 = []
# cifar100_convnet_5 = []
# cifar100_convnet_6 = []
# cifar100_convnet_22 = []
# cifar100_convnet_55 = []
# cifar100_convnet_66 = []
# cifar100_convnet_honest = []
# for i in cifar100_convnet:
#     if i[4]=="attacker 1":
#         cifar100_convnet_1.append(i)
#     elif i[4]=="attacker 2":
#         cifar100_convnet_2.append(i)
#     elif i[4]=="attacker 5":
#         cifar100_convnet_5.append(i)
#     elif i[4]=="attacker 6":
#         cifar100_convnet_6.append(i)
#     elif i[4]=="attacker 22":
#         cifar100_convnet_22.append(i)
#     elif i[4]=="attacker 55":
#         cifar100_convnet_55.append(i)
#     elif i[4]=="attacker 66":
#         cifar100_convnet_66.append(i)
#     elif i[4]=="honest":
#         cifar100_convnet_honest.append(i)
# #endregion

# cifar100_convnet_1_x = []
# cifar100_convnet_1_y = []
# for i in cifar100_convnet_1:
#     cifar100_convnet_1_x.append(int(i[3]))
#     cifar100_convnet_1_y.append(float(i[0]))
                                
# cifar100_convnet_2_x = []
# cifar100_convnet_2_y = []
# for i in cifar100_convnet_2:
#     cifar100_convnet_2_x.append(int(i[3]))
#     cifar100_convnet_2_y.append(float(i[0]))

# # print(cifar100_convnet_2_x)
# # print(cifar100_convnet_2_y)
# # hj

# cifar100_convnet_5_x = []
# cifar100_convnet_5_y = []
# for i in cifar100_convnet_5:
#     cifar100_convnet_5_x.append(int(i[3]))
#     cifar100_convnet_5_y.append(float(i[0]))

# cifar100_convnet_6_x = []
# cifar100_convnet_6_y = []
# for i in cifar100_convnet_6:
#     cifar100_convnet_6_x.append(int(i[3]))
#     cifar100_convnet_6_y.append(float(i[0]))

# cifar100_convnet_22_x = []
# cifar100_convnet_22_y = []
# for i in cifar100_convnet_22:
#     cifar100_convnet_22_x.append(int(i[3]))
#     cifar100_convnet_22_y.append(float(i[0]))

# cifar100_convnet_55_x = []
# cifar100_convnet_55_y = []
# for i in cifar100_convnet_55:
#     cifar100_convnet_55_x.append(int(i[3]))
#     cifar100_convnet_55_y.append(float(i[0]))

# cifar100_convnet_66_x = []
# cifar100_convnet_66_y = []
# for i in cifar100_convnet_66:
#     cifar100_convnet_66_x.append(int(i[3]))
#     cifar100_convnet_66_y.append(float(i[0]))

# cifar100_convnet_honest_x = []
# cifar100_convnet_honest_y = []
# for i in cifar100_convnet_honest:
#     cifar100_convnet_honest_x.append(int(i[3]))
#     cifar100_convnet_honest_y.append(float(i[0]))


# font1 = {'family' : 'Times New Roman',
# 'weight' : 'normal',
# 'size'   : 14,
# }

# #开始画图
# # plt.yticks(np.arange(0, 3500, 500),fontproperties = 'Times New Roman', size = 12)
# plt.xticks(np.arange(0, x_max, x_em),fontproperties = 'Times New Roman', size = 12)
# # plt.ylim(0, 3500)
# plt.xlim(0, x_max)
# plt.title('ConvNet')

# legend_elements = [Line2D([0], [0], marker='o', color='#32CD32', label=a1,
#                           markerfacecolor='#32CD32', markersize=7),
#                    Line2D([0], [0], marker='s', color='#0000CD', label=a2,
#                           markerfacecolor='#0000CD', markersize=7),
#                    Line2D([0], [0], marker='x', color='red', label=a5,
#                           markerfacecolor='red', markersize=7),
#                    Line2D([0], [0], marker='^', color='#FFCC33', label=a6,
#                           markerfacecolor='#FFCC33', markersize=7),
#                    Line2D([0], [0], marker='D', color='#1A1A1A', label=a0,
#                           markerfacecolor='#1A1A1A', markersize=7),
#                    Line2D([0], [0], color='black', lw=2, linestyle=':', label='longer'),
#                    Line2D([0], [0], color='black', lw=2, linestyle='--', label='shorter'),
#                    Line2D([0], [0], color='black', lw=2, linestyle='-', label='normal'),]

# plt.plot(cifar100_convnet_1_x, cifar100_convnet_1_y, marker = 'o', color='#32CD32', label=a1,linestyle='-',linewidth=2)
# plt.plot(cifar100_convnet_2_x, cifar100_convnet_2_y, marker = 's', color='#0000CD', label=a2,linestyle='-',linewidth=2)
# plt.plot(cifar100_convnet_22_x, cifar100_convnet_22_y, marker = 's', color='#0000CD', label=a2,linestyle=':',linewidth=2)
# plt.plot(cifar100_convnet_5_x, cifar100_convnet_5_y, marker = 'x', color='red', label=a5,linestyle='-',linewidth=2)
# plt.plot(cifar100_convnet_55_x, cifar100_convnet_55_y, marker = 'x', color='red', label=a5,linestyle='--',linewidth=2)
# plt.plot(cifar100_convnet_6_x, cifar100_convnet_6_y, marker = '^', color='#FFCC33', label=a6,linestyle='-',linewidth=2)
# plt.plot(cifar100_convnet_66_x, cifar100_convnet_66_y, marker = '^', color='#FFCC33', label=a6,linestyle='--',linewidth=2)
# plt.plot(cifar100_convnet_honest_x, cifar100_convnet_honest_y, marker = 'D', color='#1A1A1A', label=a0,linestyle='-',linewidth=2)

# #plt.plot(sub_axix, test_acys, color='red', label='testing accuracy')

# #plt.plot(x_axix, C,  color='#32CD32', label='Normalized Cost ($C\\times p/60$)',linestyle='--',linewidth=3)
# #plt.plot(x_axix, T, color='#0000CD', label='Normalized Throughput ($T_2/20$)',linewidth=3)
# plt.legend(handles=legend_elements, prop=font1,loc = 'upper center',bbox_to_anchor=(0.78,1)) # 显示图例

# plt.xlabel('Number of Images per Class (IPC)',size='16')
# plt.ylabel('Average Accuracy',size='16')
# plt.show()
# # fig.savefig("result2.pdf",dpi=1200)


# #endregion


# # region cifar100 convnetd2
# cifar100_convnetd2_1 = []
# cifar100_convnetd2_2 = []
# cifar100_convnetd2_5 = []
# cifar100_convnetd2_6 = []
# cifar100_convnetd2_22 = []
# cifar100_convnetd2_55 = []
# cifar100_convnetd2_66 = []
# cifar100_convnetd2_honest = []
# for i in cifar100_convnetd2:
#     if i[4]=="attacker 1":
#         cifar100_convnetd2_1.append(i)
#     elif i[4]=="attacker 2":
#         cifar100_convnetd2_2.append(i)
#     elif i[4]=="attacker 5":
#         cifar100_convnetd2_5.append(i)
#     elif i[4]=="attacker 6":
#         cifar100_convnetd2_6.append(i)
#     elif i[4]=="attacker 22":
#         cifar100_convnetd2_22.append(i)
#     elif i[4]=="attacker 55":
#         cifar100_convnetd2_55.append(i)
#     elif i[4]=="attacker 66":
#         cifar100_convnetd2_66.append(i)
#     elif i[4]=="honest":
#         cifar100_convnetd2_honest.append(i)
# # endregion

# cifar100_convnetd2_1_x = []
# cifar100_convnetd2_1_y = []
# for i in cifar100_convnetd2_1:
#     cifar100_convnetd2_1_x.append(int(i[3]))
#     cifar100_convnetd2_1_y.append(float(i[0]))
                                
# cifar100_convnetd2_2_x = []
# cifar100_convnetd2_2_y = []
# for i in cifar100_convnetd2_2:
#     cifar100_convnetd2_2_x.append(int(i[3]))
#     cifar100_convnetd2_2_y.append(float(i[0]))

# cifar100_convnetd2_5_x = []
# cifar100_convnetd2_5_y = []
# for i in cifar100_convnetd2_5:
#     cifar100_convnetd2_5_x.append(int(i[3]))
#     cifar100_convnetd2_5_y.append(float(i[0]))

# cifar100_convnetd2_6_x = []
# cifar100_convnetd2_6_y = []
# for i in cifar100_convnetd2_6:
#     cifar100_convnetd2_6_x.append(int(i[3]))
#     cifar100_convnetd2_6_y.append(float(i[0]))

# cifar100_convnetd2_22_x = []
# cifar100_convnetd2_22_y = []
# for i in cifar100_convnetd2_22:
#     cifar100_convnetd2_22_x.append(int(i[3]))
#     cifar100_convnetd2_22_y.append(float(i[0]))

# cifar100_convnetd2_55_x = []
# cifar100_convnetd2_55_y = []
# for i in cifar100_convnetd2_55:
#     cifar100_convnetd2_55_x.append(int(i[3]))
#     cifar100_convnetd2_55_y.append(float(i[0]))

# cifar100_convnetd2_66_x = []
# cifar100_convnetd2_66_y = []
# for i in cifar100_convnetd2_66:
#     cifar100_convnetd2_66_x.append(int(i[3]))
#     cifar100_convnetd2_66_y.append(float(i[0]))

# cifar100_convnetd2_honest_x = []
# cifar100_convnetd2_honest_y = []
# for i in cifar100_convnetd2_honest:
#     cifar100_convnetd2_honest_x.append(int(i[3]))
#     cifar100_convnetd2_honest_y.append(float(i[0]))


# font1 = {'family' : 'Times New Roman',
# 'weight' : 'normal',
# 'size'   : 14,
# }

# #开始画图
# # plt.yticks(np.arange(0, 3500, 500),fontproperties = 'Times New Roman', size = 12)
# plt.xticks(np.arange(0, x_max, x_em),fontproperties = 'Times New Roman', size = 12)
# # plt.ylim(0, 3500)
# plt.xlim(0, x_max)
# plt.title('ConvNetD2')

# legend_elements = [Line2D([0], [0], marker='o', color='#32CD32', label=a1,
#                           markerfacecolor='#32CD32', markersize=7),
#                    Line2D([0], [0], marker='s', color='#0000CD', label=a2,
#                           markerfacecolor='#0000CD', markersize=7),
#                    Line2D([0], [0], marker='x', color='red', label=a5,
#                           markerfacecolor='red', markersize=7),
#                    Line2D([0], [0], marker='^', color='#FFCC33', label=a6,
#                           markerfacecolor='#FFCC33', markersize=7),
#                    Line2D([0], [0], marker='D', color='#1A1A1A', label=a0,
#                           markerfacecolor='#1A1A1A', markersize=7),
#                    Line2D([0], [0], color='black', lw=2, linestyle=':', label='longer'),
#                    Line2D([0], [0], color='black', lw=2, linestyle='--', label='shorter'),
#                    Line2D([0], [0], color='black', lw=2, linestyle='-', label='normal'),]

# plt.plot(cifar100_convnetd2_1_x, cifar100_convnetd2_1_y, marker = 'o', color='#32CD32', label=a1,linestyle='-',linewidth=2)
# plt.plot(cifar100_convnetd2_2_x, cifar100_convnetd2_2_y, marker = 's', color='#0000CD', label=a2,linestyle='-',linewidth=2)
# plt.plot(cifar100_convnetd2_22_x, cifar100_convnetd2_22_y, marker = 's', color='#0000CD', label=a2,linestyle=':',linewidth=2)
# plt.plot(cifar100_convnetd2_5_x, cifar100_convnetd2_5_y, marker = 'x', color='red', label=a5,linestyle='-',linewidth=2)
# plt.plot(cifar100_convnetd2_55_x, cifar100_convnetd2_55_y, marker = 'x', color='red', label=a5,linestyle='--',linewidth=2)
# plt.plot(cifar100_convnetd2_6_x, cifar100_convnetd2_6_y, marker = '^', color='#FFCC33', label=a6,linestyle='-',linewidth=2)
# plt.plot(cifar100_convnetd2_66_x, cifar100_convnetd2_66_y, marker = '^', color='#FFCC33', label=a6,linestyle='--',linewidth=2)
# plt.plot(cifar100_convnetd2_honest_x, cifar100_convnetd2_honest_y, marker = 'D', color='#1A1A1A', label=a0,linestyle='-',linewidth=2)

# #plt.plot(sub_axix, test_acys, color='red', label='testing accuracy')

# #plt.plot(x_axix, C,  color='#32CD32', label='Normalized Cost ($C\\times p/60$)',linestyle='--',linewidth=3)
# #plt.plot(x_axix, T, color='#0000CD', label='Normalized Throughput ($T_2/20$)',linewidth=3)
# plt.legend(handles=legend_elements, prop=font1,loc = 'upper center',bbox_to_anchor=(0.78,1)) # 显示图例

# plt.xlabel('Number of Images per Class (IPC)',size='16')
# plt.ylabel('Average Accuracy',size='16')
# plt.show()
# # fig.savefig("result2.pdf",dpi=1200)


# # endregion
   









