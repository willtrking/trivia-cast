//
//  MediaView.swift
//  TriviaCast
//
//  Created by Brophy on 2/7/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import AVKit
import Foundation
import SnapKit
import MediaPlayer

class MediaView: UIView {
    
    override var contentMode: UIView.ContentMode {
        didSet {
            switch contentMode {
            case .scaleAspectFill:
                playerLayer.videoGravity = .resizeAspectFill
            case .scaleAspectFit:
                playerLayer.videoGravity = .resizeAspect
            default:
                playerLayer.videoGravity = .resizeAspect
            }
            
        }
    }
    
    private let player = AVQueuePlayer()
    let playerLayer: AVPlayerLayer
    private var playerItem: AVPlayerItem?
    private static var playerLooperContext = 0 // Key-value observing context
    private static var playerLayerContext = 1
    private var playerLooper: AVPlayerLooper?
    
    init() {
        self.playerLayer = AVPlayerLayer(player: player)
        super.init(frame: CGRect.zero)
        build()
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
    deinit {
        playerLayer.removeObserver(self, forKeyPath: "readyForDisplay", context: &MediaView.playerLayerContext)
        removePlayerLooperObserver()
        NotificationCenter.default.removeObserver(self)
    }
    
    func play() {
        player.play()
    }
    
    func pause() {
        player.pause()
    }
    
    func clearVideo() {
        player.pause()
        playerLooper?.disableLooping()
        player.removeAllItems()
        playerLayer.isHidden = true
        removePlayerLooperObserver()
    }
    
    func setContent(_ url: URL?) {
        guard Thread.isMainThread else {
            fatalError("You must call setContent on the main thread")
        }
        clearVideo()
        
        guard let url = url else {
            return
        }
        
        let requiredAssetKeys = [
            "playable",
            "hasProtectedContent"
        ]
        
        let asset = AVURLAsset(url: url)
        let playerItem = AVPlayerItem(asset: asset, automaticallyLoadedAssetKeys: requiredAssetKeys)
    
        player.removeAllItems()
        player.seek(to: CMTime.zero)
        
        self.playerLooper = AVPlayerLooper(player: self.player, templateItem: playerItem)
        playerLooper?.addObserver(
            self,
            forKeyPath: #keyPath(AVPlayerLooper.status),
            options: [.old, .new],
            context: &MediaView.playerLooperContext
        )
        if let status = playerLooper?.status, status == .ready {
            play()
        }
    }
    
    override func layoutSubviews() {
        super.layoutSubviews()
        playerLayer.frame = bounds
    }
    
    override func observeValue(forKeyPath keyPath: String?,
                               of object: Any?,
                               change: [NSKeyValueChangeKey: Any]?,
                               context: UnsafeMutableRawPointer?) {
        
        guard context == &MediaView.playerLooperContext || context == &MediaView.playerLayerContext else {
            super.observeValue(
                forKeyPath: keyPath,
                of: object,
                change: change,
                context: context
            )
            return
        }
        
        if keyPath == "readyForDisplay" {
            if playerLayer.isReadyForDisplay {
                playerLayer.isHidden = false
            }
        } else if keyPath == #keyPath(AVPlayerLooper.status) {
            let status: AVPlayerLooper.Status
            
            // Get the status change from the change dictionary
            if let statusNumber = change?[.newKey] as? NSNumber {
                status = AVPlayerLooper.Status(rawValue: statusNumber.intValue)!
            } else {
                status = .unknown
            }
            
            // Switch over the status
            switch status {
            case .ready:
                play()
            case .failed:
                break
            case .unknown, .cancelled:
                break
            }
        }
    }
    
    private func build() {
        self.contentMode = .scaleAspectFill
        self.clipsToBounds = true
        
        playerLayer.backgroundColor = UIColor.clear.cgColor
        self.layer.addSublayer(playerLayer)
        
        playerLayer.addObserver(
            self,
            forKeyPath: "readyForDisplay",
            options: [.new, .initial],
            context: &MediaView.playerLayerContext
        )
    }
    
    private func removePlayerLooperObserver() {
        playerLooper?.removeObserver(
            self,
            forKeyPath: #keyPath(AVPlayerLooper.status),
            context: &MediaView.playerLooperContext
        )
        // Attempting to remove the observer 2 times is illegal, so nil it out just
        // in case
        playerLooper = nil
    }
}
