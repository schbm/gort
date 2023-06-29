import matplotlib.pyplot as plt

xdata = [100, 1000, 10000, 100000, 200000, 300000]
xdataDetail = [100, 1000, 10000, 100000, 200000, 300000, 1000000, 10000000, 50000000, 100000000]
ydataQuick = [0, 0, 0, 6, 9, 16]
ydataQuickDetail = [0, 0, 0, 4, 11, 16, 43, 394, 2092, 5273]
ydataQuickOptDetail = [0, 0, 0, 1, 2, 4, 16, 176, 1026, 2081]

ydataBubble = [0, 1, 66, 6706, 39406, 75175]
ydataInsert = [0, 1, 58, 4408, 16467, 41689]
ydataSelection = [0, 0, 75, 5768, 26564, 58935]
ydataShell = [0, 0, 0, 0, 14, 14]
ydataShellDetail = [0, 0, 3, 2, 5, 6, 28, 354, 1806, 4000]

fig, axs = plt.subplots(2)
fig.suptitle('very accurate cool')
axs[0].plot(xdata, ydataSelection, color='grey', label='Selection')
axs[0].plot(xdata, ydataInsert, color='black', label='Insert')
axs[0].plot(xdata, ydataBubble, color='yellow', label='Bubble')
axs[0].plot(xdata, ydataQuick, color='red', label='Quick')
axs[0].plot(xdata, ydataShell, color='blue', label='Shell')
axs[0].set_title('empiric time dev')
axs[0].set_ylabel('time (ms)')
axs[0].set_xlabel('n')
axs[0].legend()
#axs[0].set_yscale('log')

axs[1].plot(xdataDetail, ydataQuickDetail, color='red', label='Quick')
axs[1].plot(xdataDetail, ydataQuickOptDetail, color='green', label='QuickOpt')
axs[1].plot(xdataDetail, ydataShellDetail, color='blue', label='Shell')
axs[1].set_ylabel('time')
axs[1].set_xlabel('n')
axs[1].legend()
#axs[1].set_yscale('log')

plt.show()
